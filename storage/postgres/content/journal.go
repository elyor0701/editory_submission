package content

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/content_service"
	"editory_submission/pkg/helper"
	"editory_submission/storage"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
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

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += ` AND ((title ILIKE '%' || :search || '%')
					OR (description ILIKE '%' || :search || '%')
					OR (isbn ILIKE '%' || :search || '%'))`
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

	q := query + filter + offset + limit

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := s.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
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
                         type
            	) VALUES (
            	          $1,
            	          $2,
            	          $3
            	) ON CONFLICT ON CONSTRAINT unique_journal_id_type DO
            	UPDATE SET
            	        	text = $2
				RETURNING 
							journal_id,
							text,
							type`

	err := s.db.QueryRow(ctx, query,
		in.GetJournalId(),
		in.GetText(),
		in.GetType(),
	).Scan(
		&res.JournalId,
		&res.Text,
		&res.Type,
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
					type
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
		)
		if err != nil {
			continue
		}

		res = append(res, obj)
	}

	return res, nil
}
