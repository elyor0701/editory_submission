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
        status,
        step,
        group_id,
        conflict,
        availability,
        funding,
        draft_step
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		$11,
		$12,
		$13
	) RETURNING 
		id,
		journal_id,
		type,
		title,
		author_id,
		description,
		status,
	    step,
        group_id,
        conflict,
        COALESCE(availability, '') as availability,
        COALESCE(funding, '') as funding,
		COALESCE(editor_status::VARCHAR, '') as editor_status,
		COALESCE(reviewer_status::VARCHAR, '') as reviewer_status,
		COALESCE(draft_step, '') as draft_step,
		TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	if !util.IsValidUUID(req.GroupId) {
		groupId, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}

		req.GroupId = groupId.String()
	}

	err = s.db.QueryRow(ctx, query,
		id.String(),
		req.GetJournalId(),
		req.GetType(),
		req.GetTitle(),
		req.GetAuthorId(),
		req.GetDescription(),
		req.GetStatus(),
		req.GetStep(),
		req.GetGroupId(),
		req.GetConflict(),
		req.GetAvailability(),
		req.GetFunding(),
		req.GetDraftStep(),
	).Scan(
		&res.Id,
		&res.JournalId,
		&res.Type,
		&res.Title,
		&res.AuthorId,
		&res.Description,
		&res.Status,
		&res.Step,
		&res.GroupId,
		&res.Conflict,
		&res.Availability,
		&res.Funding,
		&res.EditorStatus,
		&res.ReviewerStatus,
		&res.DraftStep,
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
	journal := &pb.Journal{}

	query := `SELECT
		d.id,
		d.journal_id,
		d.type,
		d.title,
		d.author_id,
		d.description,
		d.status,
	    d.step,
        d.group_id,
        d.conflict,
        COALESCE(availability, '') as availability,
        COALESCE(funding, '') as funding,
		COALESCE(editor_status::VARCHAR, '') as editor_status,
		COALESCE(reviewer_status::VARCHAR, '') as reviewer_status,
		COALESCE(draft_step, '') as draft_step,
		TO_CHAR(d.created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(d.updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at,
		j.id,
		j.title
	FROM
		"draft" d
	INNER JOIN "journal" j on d.journal_id = j.id
	WHERE
		d.id = $1`

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
		&res.Step,
		&res.GroupId,
		&res.Conflict,
		&res.Availability,
		&res.Funding,
		&res.EditorStatus,
		&res.ReviewerStatus,
		&res.DraftStep,
		&res.CreatedAt,
		&res.UpdatedAt,
		&journal.Id,
		&journal.Title,
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

	validArticleStatus := map[string]bool{
		config.ARTICLE_STATUS_DRAFT:                        true,
		config.ARTICLE_STATUS_NEW:                          true,
		config.ARTICLE_STATUS_PENDING:                      true,
		config.ARTICLE_STATUS_DENIED:                       true,
		config.ARTICLE_STATUS_CONFIRMED:                    true,
		config.ARTICLE_STATUS_PUBLISHED:                    true,
		config.ARTICLE_REVIEWER_STATUS_BACK_FOR_CORRECTION: true,
	}

	query := `SELECT
		d.id,
		d.journal_id,
		d.type,
		d.title,
		d.author_id,
		d.description,
		d.status,
	    d.step,
        d.group_id,
        d.conflict,
        COALESCE(availability, '') as availability,
        COALESCE(funding, '') as funding,
		COALESCE(editor_status::VARCHAR, '') as editor_status,
		COALESCE(reviewer_status::VARCHAR, '') as reviewer_status,
		COALESCE(draft_step, '') as draft_step,
		TO_CHAR(d.created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(d.updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at,
		j.id,
		j.title
	FROM
		"draft" d
	INNER JOIN "journal" j on d.journal_id = j.id`

	filter := ` WHERE 1=1`
	group := ``
	order := ` ORDER BY created_at DESC`

	if util.IsValidUUID(req.GetJournalId()) {
		filter += " AND journal_id = :journal_id"
		params["journal_id"] = req.GetJournalId()
	}

	if util.IsValidUUID(req.GetAuthorId()) {
		filter += " AND d.author_id = :author_id"
		params["author_id"] = req.GetAuthorId()
	}

	if _, ok := validArticleStatus[req.GetStatus()]; ok {
		filter += " AND d.status = :status"
		params["status"] = req.GetStatus()
	}

	if util.IsValidUUID(req.GroupId) {
		filter += " AND group_id = :group_id"
		params["group_id"] = req.GroupId
	} else {
		query += ` INNER JOIN (
						SELECT DISTINCT
							FIRST_VALUE(id) OVER (PARTITION BY group_id ORDER BY created_at DESC) AS id
						FROM "draft"
					) d2 ON d.id = d2.id`
		group = ` GROUP BY group_id`
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

	cQ := `SELECT count(1) over() FROM "draft" d` + filter + group

	cQ, arr = helper.ReplaceQueryParams(cQ, params)

	fmt.Println(cQ)

	err = s.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	q := query + filter + order + offset + limit

	q, arr = helper.ReplaceQueryParams(q, params)

	fmt.Println(q)

	rows, err := s.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &pb.Article{}
		journal := &pb.Journal{}

		err = rows.Scan(
			&obj.Id,
			&obj.JournalId,
			&obj.Type,
			&obj.Title,
			&obj.AuthorId,
			&obj.Description,
			&obj.Status,
			&obj.Step,
			&obj.GroupId,
			&obj.Conflict,
			&obj.Availability,
			&obj.Funding,
			&obj.EditorStatus,
			&obj.ReviewerStatus,
			&obj.DraftStep,
			&obj.CreatedAt,
			&obj.UpdatedAt,
			&journal.Id,
			&journal.Title,
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
		config.ARTICLE_STATUS_DRAFT:                        true,
		config.ARTICLE_STATUS_NEW:                          true,
		config.ARTICLE_STATUS_PENDING:                      true,
		config.ARTICLE_STATUS_DENIED:                       true,
		config.ARTICLE_STATUS_CONFIRMED:                    true,
		config.ARTICLE_STATUS_PUBLISHED:                    true,
		config.ARTICLE_REVIEWER_STATUS_BACK_FOR_CORRECTION: true,
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

	if req.GetDraftStep() != "" {
		querySet += `, draft_step = :draft_step`
		params["draft_step"] = req.GetDraftStep()
	}

	if util.IsValidUUID(req.GetAuthorId()) {
		querySet += `, author_id = :author_id`
		params["author_id"] = req.GetAuthorId()
	}

	if util.IsValidUUID(req.GroupId) {
		querySet += `, group_id = :group_id`
		params["group_id"] = req.GroupId
	}

	if req.GetDescription() != "" {
		querySet += `, description = :description`
		params["description"] = req.GetDescription()
	}

	if req.Step != "" {
		querySet += `, step = :step`
		params["step"] = req.Step
	}

	if _, ok := validArticleStatus[req.Status]; ok {
		querySet += `, status = :status`
		params["status"] = req.Status
	}

	if req.EditorStatus != "" {
		querySet += `, editor_status = :editor_status`
		params["editor_status"] = req.EditorStatus
	}

	if req.ReviewerStatus != "" {
		querySet += `, reviewer_status = :reviewer_status`
		params["reviewer_status"] = req.ReviewerStatus
	}

	if req.Conflict {
		querySet += `, conflict = :conflict`
		params["conflict"] = req.Conflict
	}

	if req.Availability != "" {
		querySet += `, availability = :availability`
		params["availability"] = req.Availability
	}

	if req.Funding != "" {
		querySet += `, funding = :funding`
		params["funding"] = req.Funding
	}

	query := querySet + filter
	q, arr := helper.ReplaceQueryParams(query, params)

	result, err := s.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), err
}
func (s *ArticleRepo) Delete(ctx context.Context, req *pb.DeleteArticleReq) (rowsAffected int64, err error) {
	queryArticleDelete := `DELETE FROM "draft" WHERE id = $1`

	result, err := s.db.Exec(ctx, queryArticleDelete, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
