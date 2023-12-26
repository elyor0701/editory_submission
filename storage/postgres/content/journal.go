package content

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/content_service"
	"editory_submission/pkg/helper"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"editory_submission/storage/postgres/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type JournalRepo struct {
	db *pgxpool.Pool
}

func NewJournalRepo(db *pgxpool.Pool) storage.JournalRepoI {
	return &JournalRepo{
		db: db,
	}
}

func (s *JournalRepo) Create(ctx context.Context, req *pb.CreateJournalReq) (res *pb.Journal, err error) {

	res = &pb.Journal{}

	query := `INSERT INTO "journal" (
		id,                 
    	cover_photo,           
    	title,         
    	access,        
    	description,              
    	price,        
    	isbn,              
    	author_id,
    	status
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9
	) RETURNING 
	    id, 
	    cover_photo, 
	    title, 
	    access, 
	    description, 
	    price, 
	    isbn, 
	    author_id,
	    status,
	    TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at, 
	    TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow(ctx, query,
		id.String(),
		req.GetCoverPhoto(),
		req.GetTitle(),
		req.GetAccess(),
		req.GetDescription(),
		req.GetPrice(),
		req.GetIsbn(),
		req.GetAuthorId(),
		req.GetStatus(),
	).Scan(
		&res.Id,
		&res.CoverPhoto,
		&res.Title,
		&res.Access,
		&res.Description,
		&res.Price,
		&res.Isbn,
		&res.AuthorId,
		&res.Status,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *JournalRepo) Get(ctx context.Context, req *pb.PrimaryKey) (res *pb.Journal, err error) {
	res = &pb.Journal{}
	author := &pb.Journal_Author{}

	query := `SELECT
		j.id,                 
    	j.cover_photo,           
    	j.title,         
    	j.access,        
    	j.description,              
    	j.price,        
    	j.isbn,              
    	coalesce(j.author_id::VARCHAR, ''),
    	j.status,
		coalesce(j.acceptance_rate, ''),
		coalesce(j.submission_to_final_decision, ''),
		coalesce(j.acceptance_to_publication, ''),
		coalesce(j.citation_indicator, ''),
		coalesce(j.impact_factor, ''),
		coalesce(j.short_description, ''),
    	TO_CHAR(j.created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at, 
	    TO_CHAR(j.updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at,
		coalesce(u.id::VARCHAR, ''),
		coalesce(u.email::VARCHAR, ''),
		coalesce(u.first_name, ''),
		coalesce(u.last_name, '')
	FROM
		"journal" j
	LEFT JOIN "user" u ON (j.author_id is not null AND j.author_id = u.id)
	WHERE
		j.id = $1`

	err = s.db.QueryRow(
		ctx,
		query,
		req.GetId()).Scan(
		&res.Id,
		&res.CoverPhoto,
		&res.Title,
		&res.Access,
		&res.Description,
		&res.Price,
		&res.Isbn,
		&res.AuthorId,
		&res.Status,
		&res.AcceptanceRate,
		&res.SubmissionToFinalDecision,
		&res.AcceptanceToPublication,
		&res.CitationIndicator,
		&res.ImpactFactor,
		&res.ShortDescription,
		&res.CreatedAt,
		&res.UpdatedAt,
		&author.Id,
		&author.Email,
		&author.FirstName,
		&author.LastName,
	)

	if err != nil {
		return nil, err
	}

	res.Author = author

	return res, nil
}

func (s *JournalRepo) GetList(ctx context.Context, req *pb.GetList) (res *pb.GetJournalListRes, err error) {
	res = &pb.GetJournalListRes{}
	params := make(map[string]interface{})
	var arr []interface{}
	validSortColumns := map[string]bool{
		"id":         true,
		"title":      true,
		"price":      true,
		"isbn":       true,
		"status":     true,
		"created_at": true,
	}

	query := `SELECT
		j.id,                 
    	j.cover_photo,           
    	j.title,         
    	j.access,        
    	j.description,              
    	j.price,        
    	j.isbn,              
    	coalesce(j.author_id::VARCHAR, ''),
    	j.status,
		coalesce(j.acceptance_rate, ''),
		coalesce(j.submission_to_final_decision, ''),
		coalesce(j.acceptance_to_publication, ''),
		coalesce(j.citation_indicator, ''),
		coalesce(j.impact_factor, ''),
		coalesce(j.short_description, ''),
    	TO_CHAR(j.created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at, 
	    TO_CHAR(j.updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at,
		coalesce(u.id::VARCHAR, ''),
		coalesce(u.email::VARCHAR, ''),
		coalesce(u.first_name, ''),
		coalesce(u.last_name, '')
	FROM
		"journal" j
	LEFT JOIN "user" u ON (j.author_id is not null AND j.author_id = u.id)`

	filter := " WHERE 1=1"

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	orderBy := ""

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += ` AND ((title ILIKE '%' || :search || '%')
					OR (description ILIKE '%' || :search || '%')
					OR (isbn ILIKE '%' || :search || '%'))`
	}

	if req.GetStatus() != "" {
		params["status"] = req.Status
		filter += ` AND status = :status`
	}

	if _, err := time.Parse(time.RFC3339, req.GetDateFrom()); err == nil {
		params["date_from"] = req.DateFrom
		filter += ` AND created_at >= :date_from`
	}

	if _, err := time.Parse(time.RFC3339, req.GetDateTo()); err == nil {
		params["date_to"] = req.DateTo
		filter += ` AND created_at < :date_to`
	}

	if _, ok := validSortColumns[req.GetSort()]; ok {
		params["sort"] = req.Sort
		orderBy += ` ORDER BY :sort`
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "journal"` + filter

	cQ, arr = helper.ReplaceQueryParams(cQ, params)

	err = s.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	q := query + filter + orderBy + offset + limit

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := s.db.Query(ctx, q, arr...)
	if errors.Is(err, pgx.ErrNoRows) {
		return res, nil
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &pb.Journal{}
		author := &pb.Journal_Author{}

		err = rows.Scan(
			&obj.Id,
			&obj.CoverPhoto,
			&obj.Title,
			&obj.Access,
			&obj.Description,
			&obj.Price,
			&obj.Isbn,
			&obj.AuthorId,
			&obj.Status,
			&obj.AcceptanceRate,
			&obj.SubmissionToFinalDecision,
			&obj.AcceptanceToPublication,
			&obj.CitationIndicator,
			&obj.ImpactFactor,
			&obj.ShortDescription,
			&obj.CreatedAt,
			&obj.UpdatedAt,
			&author.Id,
			&author.Email,
			&author.FirstName,
			&author.LastName,
		)
		if err != nil {
			return res, err
		}

		res.Journals = append(res.Journals, obj)
	}

	return res, nil
}

func (s *JournalRepo) Update(ctx context.Context, req *pb.Journal) (res *pb.Journal, err error) {

	validJournalStatus := map[string]bool{
		"ACTIVE":   true,
		"INACTIVE": true,
	}

	params := make(map[string]interface{})

	querySet := `UPDATE "journal" SET                
    	updated_at = CURRENT_TIMESTAMP`

	filter := ` WHERE id = :id`
	params["id"] = req.Id

	if req.CoverPhoto != "" {
		querySet += `, cover_photo = :cover_photo`
		params["cover_photo"] = req.CoverPhoto
	}

	if req.Title != "" {
		querySet += `, title = :title`
		params["title"] = req.Title
	}

	if req.Description != "" {
		querySet += `, description = :description`
		params["description"] = req.Description
	}

	if req.Price > 0 {
		querySet += `, price = :price`
		params["price"] = req.Price
	}

	if req.Isbn != "" {
		querySet += `, isbn = :isbn`
		params["isbn"] = req.Isbn
	}

	if _, ok := validJournalStatus[req.Status]; ok {
		querySet += `, status = :status`
		params["status"] = req.Status
	}

	if req.AcceptanceRate != "" {
		querySet += `, acceptance_rate = :acceptance_rate`
		params["acceptance_rate"] = req.AcceptanceRate
	}

	if req.SubmissionToFinalDecision != "" {
		querySet += `, submission_to_final_decision = :submission_to_final_decision`
		params["submission_to_final_decision"] = req.SubmissionToFinalDecision
	}

	if req.AcceptanceToPublication != "" {
		querySet += `, acceptance_to_publication = :acceptance_to_publication`
		params["acceptance_to_publication"] = req.AcceptanceToPublication
	}

	if req.CitationIndicator != "" {
		querySet += `, citation_indicator = :citation_indicator`
		params["citation_indicator"] = req.CitationIndicator
	}

	if req.ImpactFactor != "" {
		querySet += `, impact_factor = :impact_factor`
		params["impact_factor"] = req.ImpactFactor
	}

	if util.IsValidUUID(req.AuthorId) {
		querySet += `, author_id = :author_id`
		params["author_id"] = req.AuthorId
	}

	query := querySet + filter
	q, arr := helper.ReplaceQueryParams(query, params)

	_, err = s.db.Exec(ctx, q, arr...)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func (s *JournalRepo) Delete(ctx context.Context, req *pb.PrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "journal" WHERE id = $1`

	result, err := s.db.Exec(ctx, query, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (s *JournalRepo) UpsertJournalData(ctx context.Context, in *pb.JournalData) (*pb.JournalData, error) {
	res := &pb.JournalData{}
	fmt.Println("UpsertJournalData--->", in)

	query := `INSERT INTO journal_data(
                         journal_id, 
                         text, 
                         type,
                         short_desc
            	) VALUES (
            	          $1,
            	          $2,
            	          $3,
            	          $4
            	) ON CONFLICT ON CONSTRAINT journal_data_journal_id_type_key DO
            	UPDATE SET
            	        	text = $2,
            	        	short_desc = $4
				RETURNING 
							journal_id,
							text,
							type,
							short_desc`

	err := s.db.QueryRow(ctx, query,
		in.GetJournalId(),
		in.GetText(),
		in.GetType(),
		in.GetShortText(),
	).Scan(
		&res.JournalId,
		&res.Text,
		&res.Type,
		&res.ShortText,
	)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *JournalRepo) GetJournalData(ctx context.Context, in *pb.PrimaryKey) ([]*pb.JournalData, error) {
	res := make([]*pb.JournalData, 0, 10)

	query := `select
					journal_id,
					text,
					type,
					short_desc
				from journal_data
				where journal_id = $1`

	rows, err := s.db.Query(ctx, query, in.GetId())
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &pb.JournalData{}

		err = rows.Scan(
			&obj.JournalId,
			&obj.Text,
			&obj.Type,
			&obj.ShortText,
		)
		if err != nil {
			continue
		}

		res = append(res, obj)
	}

	return res, nil
}

func (s *JournalRepo) UpsertSubject(ctx context.Context, in *models.UpsertJournalSubjectReq) (*models.UpsertJournalSubjectRes, error) {
	res := &models.UpsertJournalSubjectRes{}

	query := `INSERT INTO journal_subject(
                        id,
                        journal_id, 
						subject_id
            	) VALUES (
            	          $1,
            	          $2,
            	          $3
            	) ON CONFLICT ON CONSTRAINT journal_subject_unique_journal_subject DO NOTHING 
				RETURNING 
				    		id,
							journal_id,
							subject_id`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow(ctx, query,
		id.String(),
		in.JournalId,
		in.SubjectId,
	).Scan(
		&res.Id,
		&res.JournalId,
		&res.SubjectId,
	)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *JournalRepo) GetSubject(ctx context.Context, in *pb.PrimaryKey) ([]*pb.Subject, error) {
	res := make([]*pb.Subject, 0, 10)

	query := `select
					subject_id,
					title
				from journal_subject
				inner join subject as s on subject_id = s.id
				where journal_id = $1`

	rows, err := s.db.Query(ctx, query, in.GetId())
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &pb.Subject{}

		err = rows.Scan(
			&obj.Id,
			&obj.Title,
		)
		if err != nil {
			continue
		}

		res = append(res, obj)
	}

	return res, nil
}

func (s *JournalRepo) DeleteSubject(ctx context.Context, in *pb.PrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "journal_subject" WHERE journal_id = $1`

	result, err := s.db.Exec(ctx, query, in.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
