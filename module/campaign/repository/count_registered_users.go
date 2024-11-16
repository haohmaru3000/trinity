package campaignrepo

import "context"

type CountRegisteredUsersStore interface {
	CountRegisteredUsers(ctx context.Context, campaignId int) (int, error)
}

type countRegisteredUsersRepo struct {
	store CountRegisteredUsersStore
}

func NewCountRegisteredUsersRepo(store CountRegisteredUsersStore) *countRegisteredUsersRepo {
	return &countRegisteredUsersRepo{store: store}
}

func (repo *countRegisteredUsersRepo) CountRegisteredUsers(ctx context.Context, campaignId int) (int, error) {
	result, err := repo.store.CountRegisteredUsers(ctx, campaignId)

	if err != nil {
		return 0, err
	}

	return result, nil
}
