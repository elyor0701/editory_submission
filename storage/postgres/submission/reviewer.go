package submission

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/submission_service"
	"editory_submission/pkg/helper"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ReviewerRepo struct {
	db *pgxpool.Pool
}

func NewReviewerRepo(db *pgxpool.Pool) storage.ReviewerRepoI {
	return &ReviewerRepo{
		db: db,
	}
}

func (s *ReviewerRepo) Create(ctx context.Context, req *pb.CreateArticleReviewerReq) (res *pb.CreateArticleReviewerRes, err error) {

	res = &pb.CreateArticleReviewerRes{}

	query := `INSERT INTO "article_reviewer" (
		id,                            
    	reviewer_id,
        article_id,
        status,
        comment
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	) RETURNING 
	    id, 
	    reviewer_id,
        article_id,
        status,
        COALESCE(comment, '') as comment,
        TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
        TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow(ctx, query,
		id.String(),
		req.ReviewerId,
		req.ArticleId,
		util.NewNullString(req.Status),
		util.NewNullString(req.Comment),
	).Scan(
		&res.Id,
		&res.ReviewerId,
		&res.ArticleId,
		&res.Status,
		&res.Comment,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ReviewerRepo) Get(ctx context.Context, req *pb.GetArticleReviewerReq) (res *pb.GetArticleReviewerRes, err error) {
	res = &pb.GetArticleReviewerRes{}
	article := &pb.Article{}
	user := &pb.Reviewer{}

	query := `SELECT
		r.id, 
	    r.reviewer_id,
        r.article_id,
        COALESCE(r.status::VARCHAR, '') as status,
        COALESCE(r.comment, '') as comment,
        TO_CHAR(r.created_at, ` + config.DatabaseQueryTimeLayout + `) AS r_created_at,
        TO_CHAR(r.updated_at, ` + config.DatabaseQueryTimeLayout + `) AS r_updated_at,
		a.id,                 
    	COALESCE(a.journal_id::VARCHAR, '') AS journal_id,           
    	a.type,         
    	a.title,        
    	COALESCE(a.author_id::VARCHAR, '') AS author_id,              
    	a.description,
		a.status,
		COALESCE(a.editor_id::VARCHAR, '') AS editor_id,
		COALESCE(a.editor_comment::VARCHAR, '') AS editor_comment,
    	TO_CHAR(a.created_at, ` + config.DatabaseQueryTimeLayout + `) AS a_created_at,
    	TO_CHAR(a.updated_at, ` + config.DatabaseQueryTimeLayout + `) AS a_updated_at,
		u.id,
		COALESCE(u.first_name, ''),
		COALESCE(u.last_name, ''),
		u.email
	FROM
		"article_reviewer" r
	INNER JOIN "article" a ON r.article_id = a.id
	INNER JOIN "user" u ON r.reviewer_id = u.id
	WHERE
		r.id = $1`

	err = s.db.QueryRow(
		ctx,
		query,
		req.GetId()).Scan(
		&res.Id,
		&res.ReviewerId,
		&res.ArticleId,
		&res.Status,
		&res.Comment,
		&res.CreatedAt,
		&res.UpdatedAt,
		&article.Id,
		&article.JournalId,
		&article.Type,
		&article.Title,
		&article.AuthorId,
		&article.Description,
		&article.Status,
		&article.EditorId,
		&article.EditorComment,
		&article.CreatedAt,
		&article.UpdatedAt,
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
	)

	if err != nil {
		return res, err
	}

	res.Article = article
	res.Reviewer = user
	return res, nil
}

func (s *ReviewerRepo) GetList(ctx context.Context, req *pb.GetArticleReviewerListReq) (res *pb.GetArticleReviewerListRes, err error) {
	res = &pb.GetArticleReviewerListRes{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		r.id, 
	    r.reviewer_id,
        r.article_id,
        COALESCE(r.status::VARCHAR, '') as status,
        COALESCE(r.comment, '') as comment,
        TO_CHAR(r.created_at, ` + config.DatabaseQueryTimeLayout + `) AS r_created_at,
        TO_CHAR(r.updated_at, ` + config.DatabaseQueryTimeLayout + `) AS r_updated_at,
		a.id,                 
    	COALESCE(a.journal_id::VARCHAR, '') AS journal_id,           
    	a.type,         
    	a.title,        
    	COALESCE(a.author_id::VARCHAR, '') AS author_id,              
    	a.description,
		a.status,
		COALESCE(a.editor_id::VARCHAR, '') AS editor_id,
		COALESCE(a.editor_comment::VARCHAR, '') AS editor_comment,
    	TO_CHAR(a.created_at, ` + config.DatabaseQueryTimeLayout + `) AS a_created_at,
    	TO_CHAR(a.updated_at, ` + config.DatabaseQueryTimeLayout + `) AS a_updated_at,
		u.id,
		COALESCE(u.first_name, ''),
		COALESCE(u.last_name, ''),
		u.email
	FROM
		"article_reviewer" r
	INNER JOIN "article" a ON r.article_id = a.id
	INNER JOIN "user" u ON r.reviewer_id = u.id`
	filter := " WHERE 1=1"

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += ` AND (status ILIKE '%' || :search || '%')`
	}

	if util.IsValidUUID(req.ReviewerId) {
		params["reviewer_id"] = req.ReviewerId
		filter += ` AND reviewer_id = :reviewer_id`
	}

	if util.IsValidUUID(req.ArticleId) {
		params["article_id"] = req.ArticleId
		filter += ` AND article_id = :article_id`
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "article_reviewer"` + filter

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
		obj := &pb.GetArticleReviewerListRes_ArticleReview{}
		article := &pb.Article{}
		user := &pb.Reviewer{}

		err = rows.Scan(
			&obj.Id,
			&obj.ReviewerId,
			&obj.ArticleId,
			&obj.Status,
			&obj.Comment,
			&obj.CreatedAt,
			&obj.UpdatedAt,
			&article.Id,
			&article.JournalId,
			&article.Type,
			&article.Title,
			&article.AuthorId,
			&article.Description,
			&article.Status,
			&article.EditorId,
			&article.EditorComment,
			&article.CreatedAt,
			&article.UpdatedAt,
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
		)
		if err != nil {
			return res, err
		}

		obj.Article = article
		obj.Reviewer = user

		res.ArticleReviewers = append(res.ArticleReviewers, obj)
	}

	return res, nil
}

func (s *ReviewerRepo) Update(ctx context.Context, req *pb.UpdateArticleReviewerReq) (rowsAffected int64, err error) {

	validArticleReviewerStatus := map[string]bool{
		config.ARTICLE_REVIEWER_STATUS_PENDING:  true,
		config.ARTICLE_REVIEWER_STATUS_APPROVED: true,
		config.ARTICLE_REVIEWER_STATUS_REJECTED: true,
	}

	querySet := `UPDATE "article_reviewer" SET                
    	updated_at = CURRENT_TIMESTAMP`

	filter := ` WHERE id = :id`

	params := map[string]interface{}{
		"id": req.GetId(),
	}

	if _, ok := validArticleReviewerStatus[req.Status]; ok {
		querySet += `, status = :status`
		params["status"] = req.Status
	}

	if req.Comment != "" {
		querySet += `, comment = :comment`
		params["comment"] = req.Comment
	}

	query := querySet + filter

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := s.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
func (s *ReviewerRepo) Delete(ctx context.Context, req *pb.DeleteArticleReviewerReq) (rowsAffected int64, err error) {
	query := `DELETE FROM "article_reviewer" WHERE id = $1 OR reviewer_id = $1`

	result, err := s.db.Exec(ctx, query, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
