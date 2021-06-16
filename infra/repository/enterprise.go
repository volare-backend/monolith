// Code generated by roche

package repository

type Enterprise struct {
	DB *sql.DB
}

func NewEnterpriseRepository(db *sql.DB) repository.IEnterpriseRepository {
	return &EnterpriseRepository{DB: db}
}
func (r EnterpriseRepository) GetList() ([]*entity.Enterprise, error) {
	var e *entity.Enterprise
	var entities []*entity.Enterprise
	rows, err := r.DB.Query("SELECT `id`, `name`, `image_url`, `description`, `homepage`, `created_at`, `updated_at` FROM `enterprise`;")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			id          string
			name        string
			image_url   string
			description string
			homepage    string
			id          int64
		)
		err := rows.Scan(&id, &id, &name, &imageUrl, &description, &homepage)
		if err != nil {
			return nil, err
		}
		e := &entity.Enterprise{
			Description: description,
			Homepage:    homepage,
			Id:          id,
			ImageUrl:    imageUrl,
			Name:        name,
		}
		entities = append(entities, e)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return entities, err
}
func (r EnterpriseRepository) Find(id int64) (*entity.Enterprise, error) {
	stmt, err := r.DB.Prepare("SELECT `id`, `name`, `image_url`, `description`, `homepage`, `created_at`, `updated_at` FROM `enterprise` WHERE `id` = ?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var (
		id          string
		name        string
		image_url   string
		description string
		homepage    string
		id          int64
	)
	err := stmt.QueryRow(1).Scan(&id, &id, &name, &imageUrl, &description, &homepage)
	if err != nil {
		if err == sql.ErrNoRows {
			return &entity.Enterprise{}, nil
		}
		return nil, err
	}
	e := &entity.Enterprise{
		Description: description,
		Homepage:    homepage,
		Id:          id,
		ImageUrl:    imageUrl,
		Name:        name,
	}
	return e, err
}
func (u EnterpriseRepository) Create(e *entity.Enterprise) (*entity.Enterprise, error) {
	stmt, err := r.DB.Prepare("INSERT INTO `enterprise` (`id`, `name`, `image_url`, `description`, `homepage`) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Id, e.Name, e.ImageUrl, e.Description, e.Homepage)
	if err != nil {
		return nil, err
	}
	return e, err
}
func (u EnterpriseRepository) Update(id int64, e *entity.Enterprise) (*entity.Enterprise, error) {
	stmt, err := r.DB.Prepare("UPDATE `enterprise` SET  `name` = ? `image_url` = ? `description` = ? `homepage` = ? WHERE `id` = ?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, e.Id, e.Name, e.ImageUrl, e.Description, e.Homepage)
	if err != nil {
		return nil, err
	}
	return e, err
}
func (u EnterpriseRepository) Delete(id int64) error {
	stmt, err := r.DB.Prepare("DELETE FROM `enterprise` WHERE `id` = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return err
}
