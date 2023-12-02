package auth

import (
	"context"
	pb "editory_submission/genproto/auth_service"
	"editory_submission/pkg/helper"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"github.com/google/uuid"

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
    	post_code          
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
		$14
	)`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	_, err = s.db.Exec(ctx, query,
		id.String(),
		util.NewNullString(req.GetUsername()),
		req.GetFirstName(),
		req.GetLastName(),
		req.GetPhone(),
		util.NewNullString(req.GetExtraPhone()),
		req.GetEmail(),
		req.GetPassword(),
		util.NewNullString(req.GetCountryId()),
		util.NewNullString(req.GetCityId()),
		util.NewNullString(req.GetProfSphere()),
		util.NewNullString(req.GetDegree()),
		util.NewNullString(req.GetAddress()),
		util.NewNullString(req.GetPostCode()),
	)

	req.Id = id.String()

	return req, err
}

func (s *UserRepo) Get(ctx context.Context, req *pb.GetUserReq) (res *pb.User, err error) {
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
    	coalesce(country_id::VARCHAR, ''),         
    	coalesce(city_id::VARCHAR, ''),            
    	coalesce(prof_sphere, ''),        
    	coalesce(degree, ''),             
    	coalesce(address, ''),            
    	coalesce(post_code, '') 
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
    	coalesce(city_id::VARCHAR, '')
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
	if err != nil {
		return res, err
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
		)
		if err != nil {
			return res, err
		}

		res.Users = append(res.Users, obj)
	}

	return res, nil
}
func (s *UserRepo) Update(ctx context.Context, req *pb.User) (rowsAffected int64, err error) {
	query := `UPDATE "user" SET                
    	username = :username,           
    	first_name = :first_name,         
    	last_name = :last_name,        
    	phone = :phone,              
    	extra_phone = :extra_phone,                              
    	country_id = :country_id,         
    	city_id = :city_id,            
    	prof_sphere = :prof_sphere,        
    	degree = :degree,             
    	address = :address,            
    	post_code = :post_code  
	WHERE
		id = :id OR email = :email`

	params := map[string]interface{}{
		"username":    req.GetUsername(),
		"first_name":  req.GetFirstName(),
		"last_name":   req.GetLastName(),
		"phone":       req.GetPhone(),
		"extra_phone": util.NewNullString(req.GetExtraPhone()),
		"country_id":  util.NewNullString(req.GetCountryId()),
		"city_id":     util.NewNullString(req.GetCityId()),
		"prof_sphere": req.GetProfSphere(),
		"degree":      req.GetDegree(),
		"address":     req.GetAddress(),
		"post_code":   req.GetPostCode(),
		"id":          req.GetId(),
		"email":       req.GetEmail(),
	}

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
