package auth

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/auth_service"
	"editory_submission/pkg/helper"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"editory_submission/storage/postgres/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) storage.UserRepoI {
	return &UserRepo{
		db: db,
	}
}

func (s *UserRepo) Create(ctx context.Context, req *pb.User) (res *pb.User, err error) {

	query := `INSERT INTO "user" (
		id,                 
    	username,           
    	first_name,         
    	last_name,        
    	phone,              
    	extra_phone,        
    	email,              
    	password,           
    	country_id,         
    	city_id,            
    	prof_sphere,        
    	degree,             
    	address,            
    	post_code,
        gender,
        is_completed
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
		$13,
		$14,
		$15,
		$16
	)`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	_, err = s.db.Exec(ctx, query,
		id.String(),
		util.NewNullString(req.GetUsername()),
		util.NewNullString(req.GetFirstName()),
		util.NewNullString(req.GetLastName()),
		util.NewNullString(req.GetPhone()),
		util.NewNullString(req.GetExtraPhone()),
		req.GetEmail(),
		req.GetPassword(),
		util.NewNullString(req.GetCountryId()),
		util.NewNullString(req.GetCityId()),
		util.NewNullString(req.GetProfSphere()),
		util.NewNullString(req.GetDegree()),
		util.NewNullString(req.GetAddress()),
		util.NewNullString(req.GetPostCode()),
		util.NewNullString(req.GetGender()),
		req.GetIsCompleted(),
	)

	res = &pb.User{
		Id:                id.String(),
		Username:          req.GetUsername(),
		FirstName:         req.GetFirstName(),
		LastName:          req.GetLastName(),
		Phone:             req.GetPhone(),
		ExtraPhone:        req.GetExtraPhone(),
		Email:             req.GetEmail(),
		EmailVerification: false,
		Password:          req.GetPassword(),
		CountryId:         req.GetCountryId(),
		CityId:            req.GetCityId(),
		ProfSphere:        req.GetProfSphere(),
		Degree:            req.GetDegree(),
		Address:           req.GetAddress(),
		PostCode:          req.GetPostCode(),
		Gender:            req.GetGender(),
	}

	return res, err
}

func (s *UserRepo) Get(ctx context.Context, req *pb.GetUserReq) (res *pb.User, err error) {
	res = &pb.User{}

	query := `SELECT
		id,                 
    	coalesce(username, ''),           
    	coalesce(first_name, ''),         
    	coalesce(last_name, ''),        
    	coalesce(phone, ''),              
    	coalesce(extra_phone, ''),        
    	email,      
    	email_verification,
    	coalesce(country_id::VARCHAR, ''),         
    	coalesce(city_id::VARCHAR, ''),            
    	coalesce(prof_sphere, ''),        
    	coalesce(degree, ''),             
    	coalesce(address, ''),            
    	coalesce(post_code, ''), 
    	coalesce(gender::VARCHAR, '') 
	FROM
		"user"
	WHERE
		id = $1`

	err = s.db.QueryRow(
		ctx,
		query,
		req.GetId()).Scan(
		&res.Id,
		&res.Username,
		&res.FirstName,
		&res.LastName,
		&res.Phone,
		&res.ExtraPhone,
		&res.Email,
		&res.EmailVerification,
		&res.CountryId,
		&res.CityId,
		&res.ProfSphere,
		&res.Degree,
		&res.Address,
		&res.PostCode,
		&res.Gender,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *UserRepo) GetList(ctx context.Context, req *pb.GetUserListReq) (res *pb.GetUserListRes, err error) {
	res = &pb.GetUserListRes{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,                 
    	coalesce(username, ''),           
    	first_name,         
    	last_name,        
    	phone,                      
    	email,      
    	coalesce(country_id::VARCHAR, ''),         
    	coalesce(city_id::VARCHAR, ''),
    	coalesce(gender::VARCHAR, '')
	FROM
		"user"`

	filter := " WHERE 1=1"

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += ` AND ((username ILIKE '%' || :search || '%')
					OR (first_name ILIKE '%' || :search || '%')
					OR (last_name ILIKE '%' || :search || '%')
					OR (email ILIKE '%' || :search || '%'))`
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "user"` + filter

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
	if errors.Is(err, pgx.ErrNoRows) {
		return res, nil
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &pb.User{}

		err = rows.Scan(
			&obj.Id,
			&obj.Username,
			&obj.FirstName,
			&obj.LastName,
			&obj.Phone,
			&obj.Email,
			&obj.CountryId,
			&obj.CityId,
			&obj.Gender,
		)
		if err != nil {
			return res, err
		}

		res.Users = append(res.Users, obj)
	}

	return res, nil
}

func (s *UserRepo) GetListWithRole(ctx context.Context, req *pb.GetUserListByRoleReq) (res *pb.GetUserListByRoleRes, err error) {
	res = &pb.GetUserListByRoleRes{}

	m := make(map[string]*pb.User)

	params := make(map[string]interface{})
	var arr []interface{}

	count := `SELECT
		count(*)
	FROM
		"user" u
	INNER JOIN role r on u.id = r.user_id`

	with := `SELECT
		u.id
	FROM
		"user" u
	INNER JOIN role r on u.id = r.user_id`

	query := `SELECT
		u.id,
		coalesce(username, '') as username,
		first_name,
		last_name,
		phone,
		email,
		coalesce(country_id::VARCHAR, '') as country_id,
		coalesce(city_id::VARCHAR, '') as city_id,
		coalesce(gender::VARCHAR, '') as gender,
		r.id as role_id,
		r.role_type,
		coalesce(journal_id::VARCHAR, '') as journal_id
	FROM
		"user" u
	INNER JOIN "role" r ON u.id = r.user_id
	WHERE u.id in (
	    SELECT
	        id
	    from "user_filtered_list"
	)`

	filter := " WHERE 1=1"

	groupBy := ` GROUP BY u.id`

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += ` AND ((username ILIKE '%' || :search || '%')
					OR (first_name ILIKE '%' || :search || '%')
					OR (last_name ILIKE '%' || :search || '%')
					OR (email ILIKE '%' || :search || '%'))`
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	if req.GetRoleType() != "" {
		params["role_type"] = req.GetRoleType()
		filter += " AND role_type = :role_type"
	}

	if util.IsValidUUID(req.GetJournalId()) {
		params["journal_id"] = req.GetJournalId()
		filter += " AND journal_id = :journal_id"
	}

	cQ := count + filter + groupBy

	cQ, arr = helper.ReplaceQueryParams(cQ, params)

	err = s.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return res, nil
	} else if err != nil {
		return nil, err
	}

	uQ := fmt.Sprintf("WITH user_filtered_list AS (%s%s%s%s%s) %s", with, filter, groupBy, offset, limit, query)

	uQ, arr = helper.ReplaceQueryParams(uQ, params)

	rows, err := s.db.Query(ctx, uQ, arr...)
	if errors.Is(err, pgx.ErrNoRows) {
		return res, nil
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &pb.User{}
		role := &pb.Role{}

		err = rows.Scan(
			&user.Id,
			&user.Username,
			&user.FirstName,
			&user.LastName,
			&user.Phone,
			&user.Email,
			&user.CountryId,
			&user.CityId,
			&user.Gender,
			&role.Id,
			&role.RoleType,
			&role.JournalId,
		)
		if err != nil {
			return res, err
		}

		if _, ok := m[user.GetId()]; ok {
			val := m[user.GetId()]
			val.Role = append(val.Role, role)
			m[user.GetId()] = val
		} else {
			user.Role = append(user.Role, role)
			m[user.GetId()] = user
		}
	}

	for _, v := range m {
		res.Users = append(res.Users, v)
	}

	return res, nil
}

func (s *UserRepo) Update(ctx context.Context, req *pb.User) (rowsAffected int64, err error) {
	fieldVal := make([]string, 0)
	params := make(map[string]interface{})
	querySet := `UPDATE "user" SET`

	filter := ` WHERE id = :id OR email = :email`
	params["id"] = req.GetId()
	params["email"] = req.GetEmail()

	if req.GetUsername() != "" {
		fieldVal = append(fieldVal, ` username = :username`)
		params["username"] = req.GetUsername()
	}

	if req.GetFirstName() != "" {
		fieldVal = append(fieldVal, ` first_name = :first_name`)
		params["first_name"] = req.GetFirstName()
	}

	if req.GetLastName() != "" {
		fieldVal = append(fieldVal, ` last_name = :last_name`)
		params["last_name"] = req.GetLastName()
	}

	if req.GetPhone() != "" {
		fieldVal = append(fieldVal, ` phone = :phone`)
		params["phone"] = req.GetPhone()
	}

	if req.GetExtraPhone() != "" {
		fieldVal = append(fieldVal, ` extra_phone = :extra_phone`)
		params["extra_phone"] = req.GetExtraPhone()
	}

	if util.IsValidUUID(req.GetCountryId()) {
		fieldVal = append(fieldVal, ` country_id = :country_id`)
		params["country_id"] = req.GetCountryId()
	}

	if util.IsValidUUID(req.GetCityId()) {
		fieldVal = append(fieldVal, ` city_id = :city_id`)
		params["city_id"] = req.GetCityId()
	}

	if req.GetProfSphere() != "" {
		fieldVal = append(fieldVal, ` prof_sphere = :prof_sphere`)
		params["prof_sphere"] = req.GetProfSphere()
	}

	if req.GetDegree() != "" {
		fieldVal = append(fieldVal, ` degree = :degree`)
		params["degree"] = req.GetDegree()
	}

	if req.GetAddress() != "" {
		fieldVal = append(fieldVal, ` address = :address`)
		params["address"] = req.GetAddress()
	}

	if req.GetPostCode() != "" {
		fieldVal = append(fieldVal, ` post_code = :post_code`)
		params["post_code"] = req.GetPostCode()
	}

	if req.GetPhone() != "" {
		fieldVal = append(fieldVal, ` gender = :gender`)
		params["gender"] = req.GetGender()
	}

	querySet += strings.Join(fieldVal, ",")

	query := querySet + filter

	fmt.Println(query)

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := s.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
func (s *UserRepo) Delete(ctx context.Context, req *pb.DeleteUserReq) (rowsAffected int64, err error) {
	query := `DELETE FROM "user" WHERE id = $1`

	result, err := s.db.Exec(ctx, query, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (s *UserRepo) GetByEmail(ctx context.Context, req *pb.GetUserReq) (res *pb.User, err error) {
	res = &pb.User{}

	query := `SELECT
		id,                 
    	coalesce(username, ''),           
    	first_name,         
    	last_name,        
    	phone,              
    	coalesce(extra_phone, ''),        
    	email,      
    	email_verification,
    	password,
    	coalesce(country_id::VARCHAR, ''),         
    	coalesce(city_id::VARCHAR, ''),            
    	coalesce(prof_sphere, ''),        
    	coalesce(degree, ''),             
    	coalesce(address, ''),            
    	coalesce(post_code, '') 
	FROM
		"user"
	WHERE
		email = $1`

	err = s.db.QueryRow(
		ctx,
		query,
		req.GetEmail()).Scan(
		&res.Id,
		&res.Username,
		&res.FirstName,
		&res.LastName,
		&res.Phone,
		&res.ExtraPhone,
		&res.Email,
		&res.EmailVerification,
		&res.Password,
		&res.CountryId,
		&res.CityId,
		&res.ProfSphere,
		&res.Degree,
		&res.Address,
		&res.PostCode,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *UserRepo) CreateEmailVerification(ctx context.Context, req *models.CreateEmailVerificationReq) (res *models.CreateEmailVerificationRes, err error) {
	query := `INSERT INTO "email_verification" (
		email,
        token,
        expires_at,
        user_id
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`

	_, err = s.db.Exec(ctx, query,
		req.Email,
		req.Token,
		req.ExpiresAt,
		req.UserId,
	)

	res = &models.CreateEmailVerificationRes{
		Email:     req.Email,
		Token:     req.Token,
		ExpiresAt: req.ExpiresAt,
		UserId:    req.UserId,
	}

	return res, err
}
func (s *UserRepo) GetEmailVerificationList(ctx context.Context, req *models.GetEmailVerificationListReq) (res *models.GetEmailVerificationListRes, err error) {
	res = &models.GetEmailVerificationListRes{}

	query := `SELECT
		email,
		token,
		sent,
		user_id,
		COALESCE(TO_CHAR(expires_at, ` + config.DatabaseQueryTimeLayout + `)::TEXT, '') AS expires_at,
		COALESCE(TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `)::TEXT, '') AS created_at
	FROM
		"email_verification"
	WHERE
		email = $1`

	rows, err := s.db.Query(ctx, query, req.Email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.EmailVerification{}

		err = rows.Scan(
			&obj.Email,
			&obj.Token,
			&obj.Sent,
			&obj.UserId,
			&obj.ExpiresAt,
			&obj.CreatedAt,
		)
		if err != nil {
			continue
		}

		res.Tokens = append(res.Tokens, obj)
		res.Count++
	}

	return res, nil
}
func (s *UserRepo) DeleteEmailVerification(ctx context.Context, req *models.DeleteEmailVerificationReq) (rowsAffected int64, err error) {
	query := `DELETE FROM "email_verification" WHERE email = $1 OR expires_at < $2`

	result, err := s.db.Exec(ctx, query, req.Email, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
func (s *UserRepo) UpdateEmailVerification(ctx context.Context, req *models.UpdateEmailVerificationReq) (res *models.UpdateEmailVerificationRes, err error) {
	query := `UPDATE "email_verification" SET                
    	sent = $1
	WHERE
		email = $2 AND token = $3`

	_, err = s.db.Exec(ctx,
		query,
		req.Sent,
		req.Email,
		req.Token,
	)
	if err != nil {
		return nil, err
	}

	res = &models.UpdateEmailVerificationRes{
		Email: req.Email,
		Token: req.Token,
		Sent:  req.Sent,
	}

	return res, nil
}

func (s *UserRepo) UpdateUserEmailVerificationStatus(ctx context.Context, req *models.UpdateUserEmailVerificationStatusReq) (rowsAffected int64, err error) {
	query := `UPDATE "user" SET                
    	email_verification = $1
	WHERE
		email = $2`

	result, err := s.db.Exec(ctx,
		query,
		req.VerificationStatus,
		req.Email,
	)
	if err != nil {

		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, nil
}
