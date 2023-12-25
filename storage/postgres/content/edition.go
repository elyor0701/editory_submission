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

type EditionRepo struct {
	db *pgxpool.Pool
}

func NewEditionRepo(db *pgxpool.Pool) storage.EditionRepoI {
	return &EditionRepo{
		db: db,
	}
}

func (s *EditionRepo) Create(ctx context.Context, req *pb.CreateEditionReq) (res *pb.Edition, err error) {

	query := `INSERT INTO "edition" (
		id,                 
    	journal_id,           
    	edition,         
    	file,
        title,
        description             
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	)`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	_, err = s.db.Exec(ctx, query,
		id.String(),
		req.GetJournalId(),
		req.GetEdition(),
		req.GetFile(),
		req.GetTitle(),
		req.GetDescription(),
	)
	if err != nil {
		return nil, err
	}

	res = &pb.Edition{
		Id:          id.String(),
		JournalId:   req.GetJournalId(),
		Edition:     req.GetEdition(),
		File:        req.GetFile(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
	}

	return res, nil
}

func (s *EditionRepo) Get(ctx context.Context, req *pb.PrimaryKey) (res *pb.Edition, err error) {
	res = &pb.Edition{}

	query := `SELECT
		id,                 
    	journal_id,           
    	edition,         
    	coalesce(file, '') as file,
    	coalesce(title, '') as title,
    	coalesce(description, '') as description,
    	to_char(created_at, ` + config.DatabaseQueryTimeLayout + `) as created_at,
    	to_char(updated_at, ` + config.DatabaseQueryTimeLayout + `) as updated_at
	FROM
		"edition"
	WHERE
		id = $1`

	err = s.db.QueryRow(
		ctx,
		query,
		req.GetId()).Scan(
		&res.Id,
		&res.JournalId,
		&res.Edition,
		&res.File,
		&res.Title,
		&res.Description,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *EditionRepo) GetList(ctx context.Context, req *pb.GetEditionListReq) (res *pb.GetEditionListRes, err error) {
	res = &pb.GetEditionListRes{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,                 
    	coalesce(journal_id::VARCHAR, '') as journal_id,           
    	edition,         
    	file,
    	to_char(created_at, ` + config.DatabaseQueryTimeLayout + `) as created_at,
    	to_char(updated_at, ` + config.DatabaseQueryTimeLayout + `) as updated_at
	FROM
		"edition"`
	filter := " WHERE journal_id = :journal_id"
	params["journal_id"] = req.GetJournalId()

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += ` AND (file ILIKE '%' || :search || '%')`
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "edition"` + filter

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
		obj := &pb.Edition{}

		err = rows.Scan(
			&obj.Id,
			&obj.JournalId,
			&obj.Edition,
			&obj.File,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		)
		if err != nil {
			return res, err
		}

		res.Editions = append(res.Editions, obj)
	}

	return res, nil
}

func (s *EditionRepo) Update(ctx context.Context, req *pb.Edition) (res *pb.Edition, err error) {

	query := `UPDATE "edition" SET                        
    	edition = $1,         
    	file = $2,                                      
    	title = $3,                                      
    	description = $4,                                      
    	updated_at = CURRENT_TIMESTAMP
	WHERE
		id = $5`

	_, err = s.db.Exec(ctx,
		query,
		req.GetEdition(),
		req.GetFile(),
		req.GetTitle(),
		req.GetDescription(),
		req.GetId(),
	)
	if err != nil {
		return &pb.Edition{}, err
	}

	res = &pb.Edition{
		Id:          req.GetId(),
		JournalId:   req.GetJournalId(),
		Edition:     req.GetEdition(),
		File:        req.GetFile(),
		CreatedAt:   req.GetCreatedAt(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
	}

	return res, err
}
func (s *EditionRepo) Delete(ctx context.Context, req *pb.PrimaryKey) (rowsAffected int64, err error) {
	queryEditionDelete := `DELETE FROM "edition" WHERE id = $1`

	result, err := s.db.Exec(ctx, queryEditionDelete, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
