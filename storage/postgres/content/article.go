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
	"time"
)

type ArticleRepo struct {
	db *pgxpool.Pool
}

func NewArticleRepo(db *pgxpool.Pool) storage.ContentArticleRepoI {
	return &ArticleRepo{
		db: db,
	}
}

func (s *ArticleRepo) Create(ctx context.Context, req *pb.CreateArticleReq) (res *pb.CreateArticleRes, err error) {

	res = &pb.CreateArticleRes{}

	query := `INSERT INTO "article" (
		id,                            
    	title,
    	description,
        journal_id,
        edition,
        file,
        author,
        content
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8
	) RETURNING 
	    id,                            
    	title,
    	description,
        journal_id,
        edition,
        file,
        author,
        content,
        TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at, 
	    TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow(ctx, query,
		id.String(),
		req.GetTitle(),
		req.GetDescription(),
		req.GetJournalId(),
		req.GetEdition(),
		req.GetFile(),
		req.GetAuthor(),
		req.GetContent(),
	).Scan(
		&res.Id,
		&res.Title,
		&res.Description,
		&res.JournalId,
		&res.Edition,
		&res.File,
		&res.Author,
		&res.Content,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ArticleRepo) Get(ctx context.Context, req *pb.GetArticleReq) (res *pb.GetArticleRes, err error) {
	res = &pb.GetArticleRes{}

	query := `SELECT
		id,                            
    	title,
    	description,
        journal_id,
        edition,
        file,
        author,
        content,
        TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at, 
	    TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	FROM
		"article"
	WHERE
		id = $1`

	err = s.db.QueryRow(
		ctx,
		query,
		req.GetId()).Scan(
		&res.Id,
		&res.Title,
		&res.Description,
		&res.JournalId,
		&res.Edition,
		&res.File,
		&res.Author,
		&res.Content,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *ArticleRepo) GetList(ctx context.Context, req *pb.GetArticleListReq) (res *pb.GetArticleListRes, err error) {
	res = &pb.GetArticleListRes{}
	params := make(map[string]interface{})
	var arr []interface{}

	validSortColumns := map[string]bool{
		"id":         true,
		"title":      true,
		"created_at": true,
		"author":     true,
		"edition":    true,
	}

	query := `SELECT
		id,                            
    	title,
    	description,
        journal_id,
        edition,
        file,
        author,
        content,
        TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at, 
	    TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	FROM
		"article"`

	filter := " WHERE 1=1"

	orderBy := ""

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	if util.IsValidUUID(req.JournalId) {
		params["journal_id"] = req.JournalId
		filter += " AND journal_id = :journal_id"
	}

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += ` AND (title ILIKE '%' || :search || '%')`
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

	cQ := `SELECT count(1) FROM "article"` + filter

	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	fmt.Println(cQ)
	err = s.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	q := query + filter + orderBy + offset + limit

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
			&obj.Title,
			&obj.Description,
			&obj.JournalId,
			&obj.Edition,
			&obj.File,
			&obj.Author,
			&obj.Content,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		)
		if err != nil {
			return res, err
		}

		res.Articles = append(res.Articles, obj)
	}

	return res, nil
}

func (s *ArticleRepo) Update(ctx context.Context, req *pb.UpdateArticleReq) (res *pb.UpdateArticleRes, err error) {
	params := make(map[string]interface{})

	querySet := `UPDATE "article" SET                
    	updated_at = CURRENT_TIMESTAMP`

	filter := ` WHERE id = :id`
	params["id"] = req.Id

	if req.Title != "" {
		querySet += `, title = :title`
		params["title"] = req.Title
	}

	if req.Description != "" {
		querySet += `, description = :description`
		params["description"] = req.Description
	}

	if req.File != "" {
		querySet += `, file = :file`
		params["file"] = req.File
	}

	if req.Author != "" {
		querySet += `, author = :author`
		params["author"] = req.Author
	}

	if req.Content != "" {
		querySet += `, content = :content`
		params["content"] = req.Content
	}

	if req.Edition > 0 {
		querySet += `, edition = :edition`
		params["edition"] = req.Edition
	}

	query := querySet + filter
	q, arr := helper.ReplaceQueryParams(query, params)

	fmt.Println(q, req)

	_, err = s.db.Exec(ctx, q, arr...)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateArticleRes{
		Id: req.Id,
	}, nil
}
func (s *ArticleRepo) Delete(ctx context.Context, req *pb.DeleteArticleReq) (rowsAffected int64, err error) {
	query := `DELETE FROM "article" WHERE id = $1`

	result, err := s.db.Exec(ctx, query, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
