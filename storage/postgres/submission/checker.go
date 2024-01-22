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

type ReviewerRepo struct {
	db *pgxpool.Pool
}

func NewReviewerRepo(db *pgxpool.Pool) storage.ReviewerRepoI {
	return &ReviewerRepo{
		db: db,
	}
}

func (s *ReviewerRepo) Create(ctx context.Context, req *pb.CreateArticleCheckerReq) (res *pb.CreateArticleCheckerRes, err error) {

	res = &pb.CreateArticleCheckerRes{}

	query := `INSERT INTO "draft_checker" (
		id,                            
    	checker_id,
        draft_id,
        status,
    	type,
        comment
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	) RETURNING 
	    id, 
	    checker_id,
        draft_id,
        status,
	    type,
        COALESCE(comment, '') as comment,
        TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
        TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow(ctx, query,
		id.String(),
		req.CheckerId,
		req.ArticleId,
		req.Status,
		req.Type,
		util.NewNullString(req.Comment),
	).Scan(
		&res.Id,
		&res.CheckerId,
		&res.ArticleId,
		&res.Status,
		&res.Type,
		&res.Comment,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	queryComment := `INSERT INTO "file_comment" (
		id,                            
    	type,
        file_id,
    	draft_checker_id,
        comment
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	)`

	for _, val := range req.GetComments() {
		commentId, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}

		_, err = s.db.Exec(
			ctx,
			queryComment,
			commentId.String(),
			val.Type,
			val.FileId,
			id.String(),
			val.Comment,
		)
		if err != nil {
			fmt.Printf("Checker------>Create: %s", err.Error())
			continue
		}
	}

	return res, nil
}

func (s *ReviewerRepo) Get(ctx context.Context, req *pb.GetArticleCheckerReq) (res *pb.GetArticleCheckerRes, err error) {
	res = &pb.GetArticleCheckerRes{}
	article := &pb.Article{}
	user := &pb.Checker{}

	query := `SELECT
		r.id, 
	    r.checker_id,
        r.draft_id,
        r.status,
	    r.type,
        COALESCE(r.comment, '') as comment,
        TO_CHAR(r.created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
        TO_CHAR(r.updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at,
		a.id,                 
    	COALESCE(a.journal_id::VARCHAR, '') AS journal_id,           
    	a.type,         
    	a.title,        
    	COALESCE(a.author_id::VARCHAR, '') AS author_id,              
    	a.description,
		a.status,
		a.availability,
		a.funding,
		a.group_id,
    	TO_CHAR(a.created_at, ` + config.DatabaseQueryTimeLayout + `) AS a_created_at,
    	TO_CHAR(a.updated_at, ` + config.DatabaseQueryTimeLayout + `) AS a_updated_at,
		u.id,
		COALESCE(u.first_name, ''),
		COALESCE(u.last_name, ''),
		u.email
	FROM
		"draft_checker" r
	INNER JOIN "draft" a ON r.draft_id = a.id
	INNER JOIN "user" u ON r.checker_id = u.id
	WHERE
		r.id = $1`

	err = s.db.QueryRow(
		ctx,
		query,
		req.GetId()).Scan(
		&res.Id,
		&res.CheckerId,
		&res.ArticleId,
		&res.Status,
		&res.Type,
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
		&article.Availability,
		&article.Funding,
		&article.GroupId,
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

	queryComment := `SELECT
		c.id,                            
    	c.type,
    	file_id,
    	draft_checker_id,
    	comment,
    	TO_CHAR(c.created_at, ` + config.DatabaseQueryTimeLayout + `) AS a_created_at,
    	TO_CHAR(c.updated_at, ` + config.DatabaseQueryTimeLayout + `) AS a_updated_at,
    	url
	FROM
		"file_comment" c
	INNER JOIN "file" f on c.file_id = f.id
	WHERE
		draft_checker_id = $1`

	rows, err := s.db.Query(
		ctx,
		queryComment,
		req.GetId(),
	)
	if err != nil {
		return res, err
	}

	var comments []*pb.FileComment
	for rows.Next() {
		obj := &pb.FileComment{}
		err = rows.Scan(
			&obj.Id,
			&obj.Type,
			&obj.FileId,
			&obj.DraftCheckerId,
			&obj.Comment,
			&obj.CreatedAt,
			&obj.UpdatedAt,
			&obj.FileUrl,
		)
		if err != nil {
			fmt.Printf("-------Get File Comment-------------> %s", err.Error())
			continue
		}

		comments = append(comments, obj)
	}

	res.Comments = comments
	res.ArticleIdData = article
	res.CheckerIdData = user

	return res, nil
}

func (s *ReviewerRepo) GetList(ctx context.Context, req *pb.GetArticleCheckerListReq) (res *pb.GetArticleCheckerListRes, err error) {
	res = &pb.GetArticleCheckerListRes{}

	validCheckerStatus := map[string]bool{
		config.ARTICLE_REVIEWER_STATUS_NEW:                 true,
		config.ARTICLE_REVIEWER_STATUS_PENDING:             true,
		config.ARTICLE_REVIEWER_STATUS_REJECTED:            true,
		config.ARTICLE_REVIEWER_STATUS_APPROVED:            true,
		config.ARTICLE_REVIEWER_STATUS_BACK_FOR_CORRECTION: true,
	}

	validCheckerType := map[string]bool{
		config.EDITOR:   true,
		config.REVIEWER: true,
	}

	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		r.id, 
	    r.checker_id,
        r.draft_id,
        r.status,
	    r.type,
        COALESCE(r.comment, '') as comment,
        TO_CHAR(r.created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
        TO_CHAR(r.updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at,
		a.id,                 
    	COALESCE(a.journal_id::VARCHAR, '') AS journal_id,           
    	a.type,         
    	a.title,        
    	COALESCE(a.author_id::VARCHAR, '') AS author_id,              
    	a.description,
		a.status,
		a.availability,
		a.funding,
		a.group_id,
    	TO_CHAR(a.created_at, ` + config.DatabaseQueryTimeLayout + `) AS a_created_at,
    	TO_CHAR(a.updated_at, ` + config.DatabaseQueryTimeLayout + `) AS a_updated_at,
		u.id,
		COALESCE(u.first_name, ''),
		COALESCE(u.last_name, ''),
		u.email
	FROM
		"draft_checker" r
	INNER JOIN "draft" a ON r.draft_id = a.id
	INNER JOIN "user" u ON r.checker_id = u.id`
	filter := " WHERE 1=1"

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	orderBy := ` ORDER BY r.created_at DESC`

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += ` AND (status ILIKE '%' || :search || '%')`
	}

	if util.IsValidUUID(req.CheckerId) {
		params["checker_id"] = req.CheckerId
		filter += ` AND checker_id = :checker_id`
	}

	if util.IsValidUUID(req.ArticleId) {
		params["draft_id"] = req.ArticleId
		filter += ` AND r.draft_id = :draft_id`
	}

	if _, ok := validCheckerType[req.Type]; ok {
		params["type"] = req.Type
		filter += ` AND r.type = :type`
	}

	if _, ok := validCheckerStatus[req.Status]; ok {
		params["status"] = req.Status
		filter += ` AND r.status = :status`
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "draft_checker" r` + filter

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
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &pb.GetArticleCheckerListRes_ArticleChecker{}
		article := &pb.Article{}
		user := &pb.Checker{}

		err = rows.Scan(
			&obj.Id,
			&obj.CheckerId,
			&obj.ArticleId,
			&obj.Status,
			&obj.Type,
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
			&article.Availability,
			&article.Funding,
			&article.GroupId,
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

		obj.ArticleIdData = article
		obj.CheckerIdData = user

		res.ArticleCheckers = append(res.ArticleCheckers, obj)
	}

	return res, nil
}

func (s *ReviewerRepo) Update(ctx context.Context, req *pb.UpdateArticleCheckerReq) (rowsAffected int64, err error) {
	rowsAffected = 0
	validArticleReviewerStatus := map[string]bool{
		config.ARTICLE_REVIEWER_STATUS_NEW:                 true,
		config.ARTICLE_REVIEWER_STATUS_PENDING:             true,
		config.ARTICLE_REVIEWER_STATUS_APPROVED:            true,
		config.ARTICLE_REVIEWER_STATUS_REJECTED:            true,
		config.ARTICLE_REVIEWER_STATUS_BACK_FOR_CORRECTION: true,
	}

	querySet := `UPDATE "draft_checker" SET                
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
	rowsAffected += result.RowsAffected()

	queryCommentInsert := `INSERT INTO "file_comment" (
		id,
		type,
		file_id,
		draft_checker_id,
		comment
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	)`

	queryCommentUpdate := `UPDATE "file_comment" SET
		comment = $1,
		updated_at = CURRENT_TIMESTAMP
	WHERE id = $2`

	for _, val := range req.GetComments() {
		if util.IsValidUUID(val.Id) {
			c, err := s.db.Exec(
				ctx,
				queryCommentUpdate,
				val.Comment,
				val.Id,
			)
			if err != nil {
				fmt.Printf("Checker------>Update: %s", err.Error())
				continue
			}

			rowsAffected += c.RowsAffected()
		} else {
			commentId, err := uuid.NewRandom()
			if err != nil {
				fmt.Printf("Checker--------->Create: %s", err.Error())
				continue
			}

			c, err := s.db.Exec(
				ctx,
				queryCommentInsert,
				commentId.String(),
				val.Type,
				val.FileId,
				req.Id,
				val.Comment,
			)
			if err != nil {
				fmt.Printf("Checker------>Create: %s", err.Error())
				continue
			}

			rowsAffected += c.RowsAffected()
		}
	}

	return rowsAffected, nil
}
func (s *ReviewerRepo) Delete(ctx context.Context, req *pb.DeleteArticleCheckerReq) (rowsAffected int64, err error) {
	query := `DELETE FROM "draft_checker" WHERE id = $1 OR checker_id = $1`

	result, err := s.db.Exec(ctx, query, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
