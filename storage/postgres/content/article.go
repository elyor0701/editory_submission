package content

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/content_service"
	"editory_submission/pkg/helper"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"fmt"
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
    	type,         
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
		req.GetType(),
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
		Type:        req.GetType(),
		AuthorId:    req.GetAuthorId(),
	}

	query = `INSERT INTO "file" (
		id,                 
    	url,           
    	type,         
    	article_id          
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`

	for _, val := range req.GetFiles() {
		id, err = uuid.NewRandom()
		if err != nil {
			return nil, err
		}

		_, err = s.db.Exec(ctx, query,
			id.String(),
			val.GetUrl(),
			val.GetType(),
			res.GetId(),
		)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (s *ArticleRepo) Get(ctx context.Context, req *pb.PrimaryKey) (res *pb.Article, err error) {
	res = &pb.Article{}

	query := `SELECT
		id,                 
    	coalesce(journal_id::VARCHAR, ''),           
    	type,         
    	title,        
    	coalesce(author_id::VARCHAR, ''),              
    	description,
    	to_char(created_at, ` + config.DatabaseQueryTimeLayout + `) as created_at
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
		&res.Type,
		&res.Title,
		&res.AuthorId,
		&res.Description,
		&res.CreatedAt,
	)

	if err != nil {
		return res, err
	}

	queryFile := `SELECT
		id,                 
    	url,           
    	type
	FROM
		"file"
	WHERE
		article_id = $1`

	rows, err := s.db.Query(ctx, queryFile, res.GetId())
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		f := pb.File{}
		err = rows.Scan(
			&f.Id,
			&f.Url,
			&f.Type,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}

		res.Files = append(res.Files, &f)
	}

	return res, nil
}

func (s *ArticleRepo) GetList(ctx context.Context, req *pb.GetArticleListReq) (res *pb.GetArticleListRes, err error) {
	res = &pb.GetArticleListRes{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,                 
    	coalesce(journal_id::VARCHAR, ''),           
    	type,         
    	title,        
    	coalesce(author_id::VARCHAR, ''),              
    	description,
    	to_char(created_at, ` + config.DatabaseQueryTimeLayout + `) as created_at
	FROM
		"article"`
	filter := " WHERE journal_id = :journal_id"
	params["journal_id"] = req.GetJournalId()

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
			&obj.Type,
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

	query := `UPDATE "article" SET                
    	journal_id = $1,           
    	type = $2,         
    	title = $3,        
    	author_id = $4,              
    	description = $5,                              
    	updated_at = CURRENT_TIMESTAMP
	WHERE
		id = $6`

	_, err = s.db.Exec(ctx,
		query,
		req.GetJournalId(),
		req.GetType(),
		req.GetTitle(),
		util.NewNullString(req.GetAuthorId()),
		req.GetDescription(),
		req.GetId(),
	)
	if err != nil {
		return &pb.Article{}, err
	}

	res = &pb.Article{
		Id:          req.GetId(),
		JournalId:   req.GetJournalId(),
		Type:        req.GetType(),
		Title:       req.GetTitle(),
		AuthorId:    req.GetAuthorId(),
		Description: req.GetDescription(),
		CreatedAt:   req.GetCreatedAt(),
	}

	queryFileUpdate := `UPDATE "file" SET                
    	url = $1,           
    	type = $2
	WHERE
		id = $3`

	queryFileInsert := `INSERT INTO "file" (
		id,                 
    	url,           
    	type,         
    	article_id          
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`

	for _, val := range req.GetFiles() {
		if util.IsValidUUID(val.GetId()) {
			_, err = s.db.Exec(ctx, queryFileUpdate,
				val.GetUrl(),
				val.GetType(),
				val.GetId(),
			)
			if err != nil {
				return nil, err
			}

			res.Files = append(res.Files, &pb.File{
				Id:        val.GetId(),
				Url:       val.GetUrl(),
				Type:      val.GetUrl(),
				ArticleId: res.GetId(),
			})
		} else {
			id, err := uuid.NewRandom()
			if err != nil {
				return nil, err
			}

			_, err = s.db.Exec(ctx, queryFileInsert,
				id.String(),
				val.GetUrl(),
				val.GetType(),
				res.GetId(),
			)
			if err != nil {
				return nil, err
			}

			res.Files = append(res.Files, &pb.File{
				Id:        id.String(),
				Url:       val.GetUrl(),
				Type:      val.GetUrl(),
				ArticleId: res.GetId(),
			})
		}
	}

	return res, err
}
func (s *ArticleRepo) Delete(ctx context.Context, req *pb.PrimaryKey) (rowsAffected int64, err error) {
	queryArticleDelete := `DELETE FROM "article" WHERE id = $1`
	queryFileDelete := `DELETE FROM "file" WHERE article_id = $1 AND draft_id is NULL`
	queryFileUpdate := `UPDATE "file" SET                
    	article_id = NULL
	WHERE
		article_id = $1`

	_, err = s.db.Exec(ctx, queryFileDelete, req.GetId())
	if err != nil {
		return 0, err
	}

	_, err = s.db.Exec(ctx, queryFileUpdate, req.GetId())
	if err != nil {
		return 0, err
	}

	result, err := s.db.Exec(ctx, queryArticleDelete, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
