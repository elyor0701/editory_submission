package submission

import (
	"context"
	pb "editory_submission/genproto/submission_service"
	"editory_submission/pkg/helper"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
)

type CoAuthorRepo struct {
	db *pgxpool.Pool
}

func NewCoAuthorRepo(db *pgxpool.Pool) storage.CoAuthorRepoI {
	return &CoAuthorRepo{
		db: db,
	}
}

func (s CoAuthorRepo) Create(ctx context.Context, req *pb.AddCoAuthorReq) (res *pb.AddCoAuthorRes, err error) {
	res = &pb.AddCoAuthorRes{}

	query := `INSERT INTO "coauthor" (
		id,                            
    	article_id,
        user_id
	) VALUES (
		$1,
		$2,
		$3
	) RETURNING 
	    id,                            
    	article_id,
    	user_id`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow(ctx, query,
		id.String(),
		req.ArticleId,
		req.UserId,
	).Scan(
		&res.Id,
		&res.ArticleId,
		&res.UserId,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s CoAuthorRepo) GetList(ctx context.Context, req *pb.GetCoAuthorsReq) (res *pb.GetCoAuthorsRes, err error) {
	res = &pb.GetCoAuthorsRes{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		c.id,                            
    	c.article_id,
        c.user_id,
        u.id,
        COALESCE(u.first_name, ''),
        COALESCE(u.last_name, ''),
        u.email,
        COALESCE(u.university_id::VARCHAR, ''),
        COALESCE(u.country_id::VARCHAR, '')
	FROM
		"coauthor" c
	INNER JOIN "user" u on c.user_id = u.id`

	filter := " WHERE 1=1"

	offset := " OFFSET 0"

	limit := " LIMIT 50"

	if util.IsValidUUID(req.DraftId) {
		params["article_id"] = req.DraftId
		filter += " AND article_id = :article_id"
	}

	cQ := `SELECT count(1) FROM "coauthor" c` + filter

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
		obj := &pb.CoAuthor{}
		author := &pb.CoAuthor_Author{}

		err = rows.Scan(
			&obj.Id,
			&obj.ArticleId,
			&obj.UserId,
			&author.Id,
			&author.FirstName,
			&author.LastName,
			&author.Email,
			&author.UniversityId,
			&author.CountryId,
		)
		if err != nil {
			return res, err
		}

		obj.UserIdData = author
		res.Coauthors = append(res.Coauthors, obj)
	}

	return res, nil
}

func (s CoAuthorRepo) Delete(ctx context.Context, req *pb.DeleteCoAuthorReq) (rowsAffected int64, err error) {
	query := `DELETE FROM "coauthor" WHERE id = $1`

	ids := strings.Split(req.Ids, ",")

	for _, id := range ids {
		result, err := s.db.Exec(ctx, query, id)
		if err != nil {
			return 0, err
		}

		rowsAffected += result.RowsAffected()
	}

	return rowsAffected, err
}
