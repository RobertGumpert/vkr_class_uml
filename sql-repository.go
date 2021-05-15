package repository

import (
	"encoding/json"
	"errors"
	"github.com/RobertGumpert/vkr-pckg/dataModel"
	"github.com/RobertGumpert/vkr-pckg/runtimeinfo"
)

type SQLRepository struct {
	storage *ApplicationStorageProvider
}

func (s *SQLRepository) AddNumbersIntersections(intersections []dataModel.NumberIssueIntersectionsModel) error {
	tx := s.storage.SqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&intersections).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *SQLRepository) AddNumberIntersections(intersection *dataModel.NumberIssueIntersectionsModel) error {
	tx := s.storage.SqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(intersection).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *SQLRepository) GetNumberIntersectionsForRepository(repositoryID uint) ([]dataModel.NumberIssueIntersectionsModel, error) {
	var intersections []dataModel.NumberIssueIntersectionsModel
	if err := s.storage.SqlDB.Where("repository_id = ?", repositoryID).Find(&intersections).Error; err != nil {
		return intersections, err
	}
	return intersections, nil
}

func (s *SQLRepository) GetNumberIntersectionsForPair(repositoryID, comparableRepositoryID uint) (dataModel.NumberIssueIntersectionsModel, error) {
	var intersection dataModel.NumberIssueIntersectionsModel
	if err := s.storage.SqlDB.Where("repository_id = ? AND comparable_repository_id = ?", repositoryID, comparableRepositoryID).Find(&intersection).Error; err != nil {
		return intersection, err
	}
	return intersection, nil
}

func (s *SQLRepository) AddRepository(repository *dataModel.RepositoryModel) error {
	tx := s.storage.SqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(repository).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *SQLRepository) AddRepositories(repositories []dataModel.RepositoryModel) error {
	tx := s.storage.SqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&repositories).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *SQLRepository) AddNearestRepositories(repositoryId uint, nearest dataModel.NearestRepositoriesJSON) error {
	bts, err := json.Marshal(nearest)
	if err != nil {
		return err
	}
	tx := s.storage.SqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var model = dataModel.NearestRepositoriesModel{
		RepositoryID: repositoryId,
		Repositories: bts,
	}
	if err := tx.Create(&model).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *SQLRepository) UpdateNearestRepositories(repositoryId uint, nearest dataModel.NearestRepositoriesJSON) error {
	return nil
}

func (s *SQLRepository) GetRepositoryByName(name string) (dataModel.RepositoryModel, error) {
	var repository dataModel.RepositoryModel
	if err := s.storage.SqlDB.Where("name = ?", name).First(&repository).Error; err != nil {
		return repository, err
	}
	return repository, nil
}

func (s *SQLRepository) GetRepositoryByID(repositoryId uint) (dataModel.RepositoryModel, error) {
	var repository dataModel.RepositoryModel
	if err := s.storage.SqlDB.Where("id = ?", repositoryId).First(&repository).Error; err != nil {
		return repository, err
	}
	return repository, nil
}

func (s *SQLRepository) GetNearestRepositories(repositoryId uint) (dataModel.NearestRepositoriesJSON, error) {
	var (
		repository dataModel.NearestRepositoriesModel
		nearest    dataModel.NearestRepositoriesJSON
	)
	if err := s.storage.SqlDB.Where("repository_id = ?", repositoryId).First(&repository).Error; err != nil {
		return nearest, err
	}
	if err := json.Unmarshal(repository.Repositories, &nearest); err != nil {
		return nearest, err
	}
	return nearest, nil
}

func (s *SQLRepository) GetAllRepositories() ([]dataModel.RepositoryModel, error) {
	tx := s.storage.SqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var model []dataModel.RepositoryModel
	if err := tx.Find(&model).Error; err != nil {
		tx.Rollback()
		return model, err
	}
	return model, tx.Commit().Error
}

func (s *SQLRepository) RewriteAllNearestRepositories(repositoryId []uint, models []dataModel.NearestRepositoriesJSON) error {
	return nil
}

func (s *SQLRepository) AddIssue(issue *dataModel.IssueModel) error {
	tx := s.storage.SqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(issue).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *SQLRepository) AddIssues(issues []dataModel.IssueModel) error {
	tx := s.storage.SqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&issues).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *SQLRepository) GetNearestIssuesForPairRepositories(mainRepositoryID, secondRepositoryID uint) ([]dataModel.NearestIssuesModel, error) {
	var model []dataModel.NearestIssuesModel
	if err := s.storage.SqlDB.Where("repository_id = ? AND repository_id_nearest_issue = ?", mainRepositoryID, secondRepositoryID).Find(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}

func (s *SQLRepository) AddNearestIssues(nearest dataModel.NearestIssuesModel) error {
	tx := s.storage.SqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&nearest).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *SQLRepository) AddListNearestIssues(nearest []dataModel.NearestIssuesModel) error {
	tx := s.storage.SqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&nearest).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *SQLRepository) GetIssueByID(issueId uint) (dataModel.IssueModel, error) {
	var model dataModel.IssueModel
	if err := s.storage.SqlDB.Where("id = ?", issueId).First(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}

func (s *SQLRepository) GetIssuesOnlyGroupRepositories(repositoryId ...uint) ([]dataModel.IssueModel, error) {
	var (
		model []dataModel.IssueModel
		id    = make([]uint, 0)
	)
	id = append(id, repositoryId...)
	if err := s.storage.SqlDB.Where("repository_id IN ?", id).Find(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}

func (s *SQLRepository) GetIssuesBesidesGroupRepositories(repositoryId ...uint) ([]dataModel.IssueModel, error) {
	var (
		model []dataModel.IssueModel
		id    = make([]uint, 0)
	)
	id = append(id, repositoryId...)
	if err := s.storage.SqlDB.Where("repository_id NOT IN ?", id).Find(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}

func (s *SQLRepository) GetIssueRepository(repositoryId uint) ([]dataModel.IssueModel, error) {
	var model []dataModel.IssueModel
	if err := s.storage.SqlDB.Where("repository_id = ?", repositoryId).Find(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}

func (s *SQLRepository) GetNearestIssuesForIssue(issueId uint) ([]dataModel.NearestIssuesModel, error) {
	var model []dataModel.NearestIssuesModel
	if err := s.storage.SqlDB.Where("issue_id = ?", issueId).Find(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}

func (s *SQLRepository) GetNearestIssuesForRepository(repositoryId uint) ([]dataModel.NearestIssuesModel, error) {
	var model []dataModel.NearestIssuesModel
	if err := s.storage.SqlDB.Where("repository_id = ?", repositoryId).Find(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}

func (s *SQLRepository) AddKeyWord(keyWord string, position int64, repositories dataModel.RepositoriesIncludeKeyWordsJSON) (dataModel.RepositoriesKeyWordsModel, error) {
	var model dataModel.RepositoriesKeyWordsModel
	bts, err := json.Marshal(repositories)
	if err != nil {
		return model, err
	}
	model = dataModel.RepositoriesKeyWordsModel{
		KeyWord:      keyWord,
		Repositories: bts,
	}
	tx := s.storage.SqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&model).Error; err != nil {
		tx.Rollback()
		return model, err
	}
	return model, tx.Commit().Error
}

func (s *SQLRepository) UpdateKeyWord(keyWord string, position int64, repositories dataModel.RepositoriesIncludeKeyWordsJSON) (dataModel.RepositoriesKeyWordsModel, error) {
	var model dataModel.RepositoriesKeyWordsModel
	bts, err := json.Marshal(repositories)
	if err != nil {
		return model, err
	}
	model = dataModel.RepositoriesKeyWordsModel{
		KeyWord:      keyWord,
		Repositories: bts,
	}
	tx := s.storage.SqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Updates(&model).Error; err != nil {
		tx.Rollback()
		return model, err
	}
	return model, tx.Commit().Error
}

func (s *SQLRepository) GetKeyWord(keyWord string) (dataModel.RepositoriesKeyWordsModel, error) {
	var model dataModel.RepositoriesKeyWordsModel
	if err := s.storage.SqlDB.Where("key_word = ?", keyWord).First(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}

func (s *SQLRepository) GetAllKeyWords() ([]dataModel.RepositoriesKeyWordsModel, error) {
	return nil, nil
}

func (s *SQLRepository) RewriteAllKeyWords(models []dataModel.RepositoriesKeyWordsModel) error {
	return nil
}

func (s *SQLRepository) HasEntities() error {
	db := s.storage.SqlDB.Begin()
	entities := []interface{}{
		&dataModel.RepositoryModel{},
		&dataModel.IssueModel{},
		&dataModel.NearestIssuesModel{},
		&dataModel.NearestRepositoriesModel{},
		&dataModel.RepositoriesKeyWordsModel{},
		&dataModel.NumberIssueIntersectionsModel{},
	}
	for _, entity := range entities {
		if exist := db.Migrator().HasTable(entity); !exist {
			return errors.New("Non exist table. ")
		}
	}
	return nil
}

func (s *SQLRepository) CreateEntities() error {
	db := s.storage.SqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()
	if err := db.Migrator().CreateTable(
		&dataModel.RepositoryModel{},
		&dataModel.IssueModel{},
		&dataModel.NearestIssuesModel{},
		&dataModel.NearestRepositoriesModel{},
		&dataModel.RepositoriesKeyWordsModel{},
		&dataModel.NumberIssueIntersectionsModel{},
	); err != nil {
		db.Rollback()
		return err
	}
	return db.Commit().Error
}

func (s *SQLRepository) Migration() error {
	db := s.storage.SqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()
	if err := db.AutoMigrate(
		&dataModel.RepositoryModel{},
		&dataModel.IssueModel{},
		&dataModel.NearestIssuesModel{},
		&dataModel.NearestRepositoriesModel{},
		&dataModel.RepositoriesKeyWordsModel{},
		&dataModel.NumberIssueIntersectionsModel{},
	); err != nil {
		db.Rollback()
		return err
	}
	return db.Commit().Error
}

func (s *SQLRepository) CloseConnection() error {
	db, err := s.storage.SqlDB.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewSQLRepository(storage *ApplicationStorageProvider) *SQLRepository {
	repository := &SQLRepository{storage: storage}
	err := repository.HasEntities()
	if err != nil {
		err := repository.CreateEntities()
		if err != nil {
			runtimeinfo.LogFatal(err)
		}
	}
	err = repository.Migration()
	if err != nil {
		runtimeinfo.LogFatal(err)
	}
	return repository
}
