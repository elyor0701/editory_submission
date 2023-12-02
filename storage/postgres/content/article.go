package content

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/content_service"
	"editory_submission/pkg/helper"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ArticleRepo struct {
	db *pgxpool.Pool
}

func NewArticleRepo(db *pgxpool.Pool) storage.ArticleRepoI {
	return &ArticleRepo{
		db: db,
	}
}

func (s *ArticleRepo) Create(ctx context.Context, req *pb.CreateArticleReq) (res *pb.Article, err error) {

	query := `INSERT INTO "article" (
		id,                 
    	journal_id,           
    	article_type,         
    	title,        
    	author_id,              
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
		util.NewNullString(req.GetJournalId()),
		req.GetArticleType(),
		req.GetTitle(),
		util.NewNullString(req.GetAuthorId()),
		req.GetDescription(),
	)
	if err != nil {
		return nil, err
	}

	res = &pb.Article{
		Id:          id.String(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		JournalId:   req.GetJournalId(),
		ArticleType: req.GetArticleType(),
		AuthorId:    req.GetAuthorId(),
	}

	return res, nil
}

func (s *ArticleRepo) Get(ctx context.Context, req *pb.PrimaryKey) (res *pb.Article, err error) {
	res = &pb.Article{}

	query := `SELECT
		id,                 
    	coalesce(journal_id::VARCHAR, ''),           
    	article_type,         
    	title,        
    	coalesce(author_id::VARCHAR, ''),              
    	description,
    	to_char(created_at, ` + config.DatabaseQueryTimeLayout + `) as created_at,
	FROM
		"article"
	WHERE
		id = $1`

	err = s.db.QueryRow(
		ctx,
		query,
		req.GetId()).Scan(
		&res.Id,
		&res.JournalId,
		&res.ArticleType,
		&res.Title,
		&res.AuthorId,
		&res.Description,
		&res.CreatedAt,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *ArticleRepo) GetList(ctx context.Context, req *pb.GetList) (res *pb.GetArticleListRes, err error) {
	res = &pb.GetArticleListRes{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,                 
    	coalesce(journal_id::VARCHAR, ''),           
    	article_type,         
    	title,        
    	coalesce(author_id::VARCHAR, ''),              
    	description,
    	to_char(created_at, ` + config.DatabaseQueryTimeLayout + `) as created_at,
	FROM
		"article"`
	filter := " WHERE 1=1"

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += ` AND ((title ILIKE '%' || :search || '%')
					OR (description ILIKE '%' || :search || '%'))`
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "article"` + filter

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
		obj := &pb.Article{}

		err = rows.Scan(
			&obj.Id,
			&obj.JournalId,
			&obj.ArticleType,
			&obj.Title,
			&obj.AuthorId,
			&obj.Description,
			&obj.CreatedAt,
		)
		if err != nil {
			return res, err
		}

		res.Articles = append(res.Articles, obj)
	}

	return res, nil
}
func (s *ArticleRepo) Update(ctx context.Context, req *pb.Article) (res *pb.Article, err error) {
	// id,
	//    	journal_id,
	//    	article_type,
	//    	title,
	//    	author_id,
	//    	description
	query := `UPDATE "article" SET                
    	journal_id = :journal_id,           
    	article_type = :article_type,         
    	title = :title,        
    	author_id = :author_id,              
    	description = :description,                              
    	updated_at = CURRENT_TIMESTAMP
	WHERE
		id = :id`

	params := map[string]interface{}{
		"journal_id":   req.GetJournalId(),
		"article_type": req.GetArticleType(),
		"title":        req.GetTitle(),
		"author_id":    req.GetAuthorId(),
		"description":  req.GetDescription(),
		"id":           req.GetId(),
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	_, err = s.db.Exec(ctx, q, arr...)
	if err != nil {
		return &pb.Article{}, err
	}

	return req, err
}
func (s *ArticleRepo) Delete(ctx context.Context, req *pb.PrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "article" WHERE id = $1`

	result, err := s.db.Exec(ctx, query, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
