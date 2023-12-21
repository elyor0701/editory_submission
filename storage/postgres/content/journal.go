package content

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/content_service"
	"editory_submission/pkg/helper"
	"editory_submission/storage"
	"editory_submission/storage/postgres/models"
	"errors"
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
    	author,
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
	    author,
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
		req.GetAuthor(),
		req.GetStatus(),
	).Scan(
		&res.Id,
		&res.CoverPhoto,
		&res.Title,
		&res.Access,
		&res.Description,
		&res.Price,
		&res.Isbn,
		&res.Author,
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

	query := `SELECT
		id,                 
    	cover_photo,           
    	title,         
    	access,        
    	description,              
    	price,        
    	isbn,              
    	author,
    	status,
    	TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at, 
	    TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	FROM
		"journal"
	WHERE
		id = $1`

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
		&res.Author,
		&res.Status,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		return res, err
	}

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
		"author":     true,
		"status":     true,
		"created_at": true,
	}

	query := `SELECT
		id,                 
    	cover_photo,           
    	title,         
    	access,        
    	description,              
    	price,        
    	isbn,              
    	author,
    	status,
    	TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at, 
	    TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	FROM
		"journal"`
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

		err = rows.Scan(
			&obj.Id,
			&obj.CoverPhoto,
			&obj.Title,
			&obj.Access,
			&obj.Description,
			&obj.Price,
			&obj.Isbn,
			&obj.Author,
			&obj.Status,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		)
		if err != nil {
			return res, err
		}

		res.Journals = append(res.Journals, obj)
	}

	return res, nil
}

func (s *JournalRepo) Update(ctx context.Context, req *pb.Journal) (res *pb.Journal, err error) {
	res = &pb.Journal{}

	query := `UPDATE "journal" SET                
    	title = :title,           
    	cover_photo = :cover_photo,         
    	access = :access,        
    	description = :description,              
    	price = :price,                              
    	isbn = :isbn,         
    	author = :author,
    	status = :status,
    	updated_at = CURRENT_TIMESTAMP
	WHERE
		id = :id
	RETURNING 
		id,
		cover_photo,
		title,
		access,
		description,
		price,
		isbn,
		author,
		status,
		TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at, 
	    TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at`

	params := map[string]interface{}{
		"title":       req.GetTitle(),
		"cover_photo": req.GetCoverPhoto(),
		"access":      req.GetAccess(),
		"description": req.GetDescription(),
		"price":       req.GetPrice(),
		"isbn":        req.GetIsbn(),
		"author":      req.GetAuthor(),
		"id":          req.GetId(),
		"status":      req.GetStatus(),
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	err = s.db.QueryRow(ctx, q, arr...).Scan(
		&res.Id,
		&res.CoverPhoto,
		&res.Title,
		&res.Access,
		&res.Description,
		&res.Price,
		&res.Isbn,
		&res.Author,
		&res.Status,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
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
            	) ON CONFLICT ON CONSTRAINT unique_journal_id_type DO
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
