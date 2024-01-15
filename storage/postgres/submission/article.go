package submission

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/submission_service"
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

func (s *ArticleRepo) Create(ctx context.Context, req *pb.CreateArticleReq) (res *pb.CreateArticleRes, err error) {
	res = &pb.CreateArticleRes{}

	query := `INSERT INTO "draft" (
		id,                 
    	journal_id,           
    	type,         
    	title,        
    	author_id,              
    	description,
        status
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	) RETURNING 
		id,
		journal_id,
		type,
		title,
		author_id,
		description,
		status,
		COALESCE(editor_id::VARCHAR, '') as editor_id,
		COALESCE(editor_comment, '') as editor_comment,
		TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow(ctx, query,
		id.String(),
		req.GetJournalId(),
		req.GetType(),
		req.GetTitle(),
		req.GetAuthorId(),
		req.GetDescription(),
		req.GetStatus(),
	).Scan(
		&res.Id,
		&res.JournalId,
		&res.Type,
		&res.Title,
		&res.AuthorId,
		&res.Description,
		&res.Status,
		&res.EditorId,
		&res.EditorComment,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, err
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
	) RETURNING 
		id,
		url,
		type,
		article_id`

	for _, val := range req.GetFiles() {
		obj := pb.File{}
		id, err = uuid.NewRandom()
		if err != nil {
			return nil, err
		}

		err = s.db.QueryRow(ctx, query,
			id.String(),
			val.GetUrl(),
			val.GetType(),
			res.GetId(),
		).Scan(
			&obj.Id,
			&obj.Url,
			&obj.Type,
			&obj.ArticleId,
		)
		if err != nil {
			return nil, err
		}

		res.Files = append(res.Files, &obj)
	}

	return res, nil
}

func (s *ArticleRepo) Get(ctx context.Context, req *pb.GetArticleReq) (res *pb.GetArticleRes, err error) {
	res = &pb.GetArticleRes{}

	query := `SELECT
		id,                 
    	coalesce(journal_id::VARCHAR, '') as journal_id,           
    	type,         
    	title,        
    	coalesce(author_id::VARCHAR, '') as author_id,              
    	description,
		status,
		coalesce(editor_id::VARCHAR, '') as editor_id,
		coalesce(editor_comment::VARCHAR, '') as editor_comment,
    	to_char(created_at, ` + config.DatabaseQueryTimeLayout + `) as created_at
    	to_char(updated_at, ` + config.DatabaseQueryTimeLayout + `) as updated_at
	FROM
		"draft"
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
		&res.Status,
		&res.EditorId,
		&res.EditorComment,
		&res.CreatedAt,
		&res.UpdatedAt,
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

	validArticleStatus := map[string]bool{
		config.ARTICLE_STATUS_NEW:       true,
		config.ARTICLE_STATUS_PENDING:   true,
		config.ARTICLE_STATUS_DENIED:    true,
		config.ARTICLE_STATUS_CONFIRMED: true,
		config.ARTICLE_STATUS_PUBLISHED: true,
	}

	query := `SELECT
		id,                 
    	COALESCE(journal_id::VARCHAR, '') AS journal_id,           
    	type,         
    	title,        
    	COALESCE(author_id::VARCHAR, '') AS author_id,              
    	description,
		status,
		COALESCE(editor_id::VARCHAR, '') AS editor_id,
		COALESCE(editor_comment::VARCHAR, '') AS editor_comment,
    	TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
    	TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	FROM
		"draft"`

	filter := `WHERE 1=1`

	if util.IsValidUUID(req.GetJournalId()) {
		filter += " AND journal_id = :journal_id"
		params["journal_id"] = req.GetJournalId()
	}

	if util.IsValidUUID(req.GetAuthorId()) {
		filter += " AND author_id = :author_id"
		params["author_id"] = req.GetAuthorId()
	}

	if _, ok := validArticleStatus[req.GetStatus()]; ok {
		filter += " AND status = :status"
		params["status"] = req.GetStatus()
	}

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

	cQ := `SELECT count(1) FROM "draft"` + filter

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
			&obj.Status,
			&obj.EditorId,
			&obj.EditorComment,
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
func (s *ArticleRepo) Update(ctx context.Context, req *pb.UpdateArticleReq) (rowsAffected int64, err error) {

	validArticleStatus := map[string]bool{
		config.ARTICLE_STATUS_NEW:       true,
		config.ARTICLE_STATUS_PENDING:   true,
		config.ARTICLE_STATUS_DENIED:    true,
		config.ARTICLE_STATUS_CONFIRMED: true,
		config.ARTICLE_STATUS_PUBLISHED: true,
	}

	params := make(map[string]interface{})

	querySet := `UPDATE "draft" SET                                              
    	updated_at = CURRENT_TIMESTAMP`

	filter := ` WHERE id = :id`
	params["id"] = req.GetId()

	if util.IsValidUUID(req.GetJournalId()) {
		querySet += `, journal_id = :journal_id`
		params["journal_id"] = req.GetJournalId()
	}

	if req.GetType() != "" {
		querySet += `, type = :type`
		params["type"] = req.GetType()
	}

	if req.GetTitle() != "" {
		querySet += `, title = :title`
		params["title"] = req.GetTitle()
	}

	if util.IsValidUUID(req.GetAuthorId()) {
		querySet += `, author_id = :author_id`
		params["author_id"] = req.GetAuthorId()
	}

	if req.GetDescription() != "" {
		querySet += `, description = :description`
		params["description"] = req.GetDescription()
	}

	if _, ok := validArticleStatus[req.Status]; ok {
		querySet += `, status = :status`
		params["status"] = req.Status
	}

	if util.IsValidUUID(req.EditorId) {
		querySet += `, editor_id = :editor_id`
		params["editor_id"] = req.EditorId
	}

	if req.EditorComment != "" {
		querySet += `, editor_comment = :editor_comment`
		params["editor_comment"] = req.EditorComment
	}

	query := querySet + filter
	q, arr := helper.ReplaceQueryParams(query, params)

	result, err := s.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
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
				return 0, err
			}
		} else {
			id, err := uuid.NewRandom()
			if err != nil {
				return 0, err
			}

			_, err = s.db.Exec(ctx, queryFileInsert,
				id.String(),
				val.GetUrl(),
				val.GetType(),
				req.GetId(),
			)
			if err != nil {
				return 0, err
			}
		}
	}

	return result.RowsAffected(), err
}
func (s *ArticleRepo) Delete(ctx context.Context, req *pb.DeleteArticleReq) (rowsAffected int64, err error) {
	queryArticleDelete := `DELETE FROM "draft" WHERE id = $1`
	//queryFileDelete := `DELETE FROM "file" WHERE article_id = $1 AND draft_id is NULL`
	//queryFileUpdate := `UPDATE "file" SET
	//	article_id = NULL
	//WHERE
	//	article_id = $1`

	//_, err = s.db.Exec(ctx, queryFileDelete, req.GetId())
	//if err != nil {
	//	return 0, err
	//}
	//
	//_, err = s.db.Exec(ctx, queryFileUpdate, req.GetId())
	//if err != nil {
	//	return 0, err
	//}

	result, err := s.db.Exec(ctx, queryArticleDelete, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
