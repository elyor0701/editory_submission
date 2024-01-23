package content

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/content_service"
	"editory_submission/pkg/helper"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type JournalAuthorRepo struct {
	db *pgxpool.Pool
}

func NewJournalAuthorRepo(db *pgxpool.Pool) storage.JournalAuthorRepoI {
	return &JournalAuthorRepo{
		db: db,
	}
}

func (s *JournalAuthorRepo) Create(ctx context.Context, req *pb.CreateJournalAuthorReq) (res *pb.CreateJournalAuthorRes, err error) {

	res = &pb.CreateJournalAuthorRes{}

	query := `INSERT INTO "journal_author" (
		id,                            
    	journal_id,
        full_name,
        photo,
        email,
        university_id,
        faculty_id
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
        full_name,
        coalesce(photo, ''),
        email,
        coalesce(university_id::VARCHAR, ''),
        coalesce(faculty_id::VARCHAR, ''),
        TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at, 
	    TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow(ctx, query,
		id.String(),
		req.JournalId,
		req.FullName,
		util.NewNullString(req.Photo),
		req.Email,
		util.NewNullString(req.UniversityId),
		util.NewNullString(req.FacultyId),
	).Scan(
		&res.Id,
		&res.JournalId,
		&res.FullName,
		&res.Photo,
		&res.Email,
		&res.UniversityId,
		&res.FacultyId,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *JournalAuthorRepo) Get(ctx context.Context, req *pb.GetJournalAuthorReq) (res *pb.GetJournalAuthorRes, err error) {
	res = &pb.GetJournalAuthorRes{}
	journal := &pb.Journal{}
	university := &pb.University{}

	query := `SELECT
		a.id,                            
    	journal_id,
        full_name,
        coalesce(photo, ''),
        email,
        coalesce(university_id::VARCHAR, ''),
        coalesce(faculty_id::VARCHAR, ''),
        TO_CHAR(a.created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at, 
	    TO_CHAR(a.updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at,
		j.id,
		j.title,
		coalesce(u.id::VARCHAR, ''),
		coalesce(u.title, ''),
		coalesce(u.logo, '')
	FROM
		"journal_author" a
	INNER JOIN "journal" j on a.journal_id = j.id
	LEFT JOIN "university" u on a.university_id = u.id
	WHERE
		id = $1`

	err = s.db.QueryRow(
		ctx,
		query,
		req.GetId()).Scan(
		&res.Id,
		&res.JournalId,
		&res.FullName,
		&res.Photo,
		&res.Email,
		&res.UniversityId,
		&res.FacultyId,
		&res.CreatedAt,
		&res.UpdatedAt,
		&journal.Id,
		&journal.Title,
		&university.Id,
		&university.Title,
		&university.Logo,
	)

	if err != nil {
		return res, err
	}

	res.JournalIdData = journal
	res.UniversityIdData = university

	return res, nil
}

func (s *JournalAuthorRepo) GetList(ctx context.Context, req *pb.GetJournalAuthorListReq) (res *pb.GetJournalAuthorListRes, err error) {
	res = &pb.GetJournalAuthorListRes{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		a.id,                            
    	journal_id,
        full_name,
        coalesce(photo, ''),
        email,
        coalesce(university_id::VARCHAR, ''),
        coalesce(faculty_id::VARCHAR, ''),
        TO_CHAR(a.created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at, 
	    TO_CHAR(a.updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at,
		j.id,
		j.title,
		coalesce(u.id::VARCHAR, ''),
		coalesce(u.title, ''),
		coalesce(u.logo, '')
	FROM
		"journal_author" a
	INNER JOIN "journal" j on a.journal_id = j.id
	LEFT JOIN "university" u on a.university_id = u.id`

	filter := " WHERE 1=1"

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += ` AND ((full_name ILIKE '%' || :search || '%')
						OR (email ILIKE '%' || :search || '%'))`
	}

	if util.IsValidUUID(req.JournalId) {
		params["journal_id"] = req.JournalId
		filter += ` AND journal_id = :journal_id`
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "journal_author"` + filter

	cQ, arr = helper.ReplaceQueryParams(cQ, params)

	err = s.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return res, nil
	} else if err != nil {
		return nil, err
	}

	q := query + filter + offset + limit

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := s.db.Query(ctx, q, arr...)
	if errors.Is(err, pgx.ErrNoRows) {
		return res, nil
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &pb.JournalAuthor{}
		journal := &pb.Journal{}
		university := &pb.University{}

		err = rows.Scan(
			&obj.Id,
			&obj.JournalId,
			&obj.FullName,
			&obj.Photo,
			&obj.Email,
			&obj.UniversityId,
			&obj.FacultyId,
			&obj.CreatedAt,
			&obj.UpdatedAt,
			&journal.Id,
			&journal.Title,
			&university.Id,
			&university.Title,
			&university.Logo,
		)
		if err != nil {
			return res, err
		}

		obj.JournalIdData = journal
		obj.UniversityIdData = university

		res.Authors = append(res.Authors, obj)
	}

	return res, nil
}

func (s *JournalAuthorRepo) Update(ctx context.Context, req *pb.UpdateJournalAuthorReq) (rowsAffected int64, err error) {
	querySet := `UPDATE "draft_checker" SET                
    	updated_at = CURRENT_TIMESTAMP`

	filter := ` WHERE id = :id`

	params := map[string]interface{}{
		"id": req.GetId(),
	}

	if util.IsValidUUID(req.GetJournalId()) {
		querySet += `, journal_id = :journal_id`
		params["journal_id"] = req.GetJournalId()
	}

	if req.FullName != "" {
		querySet += `, full_name = :full_name`
		params["full_name"] = req.FullName
	}

	if req.Photo != "" {
		querySet += `, photo = :photo`
		params["photo"] = req.Photo
	}

	if req.Email != "" {
		querySet += `, email = :email`
		params["email"] = req.Email
	}

	if util.IsValidUUID(req.UniversityId) {
		querySet += `, university_id = :university_id`
		params["university_id"] = req.UniversityId
	}

	if util.IsValidUUID(req.FacultyId) {
		querySet += `, faculty_id = :faculty_id`
		params["faculty_id"] = req.FacultyId
	}

	query := querySet + filter

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := s.db.Exec(ctx, q, arr...)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
func (s *JournalAuthorRepo) Delete(ctx context.Context, req *pb.DeleteJournalAuthorReq) (rowsAffected int64, err error) {
	query := `DELETE FROM "journal_author" WHERE id = $1`

	result, err := s.db.Exec(ctx, query, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
