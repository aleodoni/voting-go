package persistence

import (
	"context"
	"errors"
	"fmt"

	"github.com/aleodoni/voting-go/internal/domain/usuario"
	mappers "github.com/aleodoni/voting-go/internal/infrastructure/persistence/mappers"
	db "github.com/aleodoni/voting-go/internal/infrastructure/persistence/sqlc/generated"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type usuarioRepositorySQLC struct {
	q *db.Queries
}

func NewUsuarioRepositorySQLC(pool *pgxpool.Pool) usuario.UsuarioRepository {
	return &usuarioRepositorySQLC{
		q: db.New(pool),
	}
}

func (r *usuarioRepositorySQLC) FindByKeycloakID(
	ctx context.Context,
	keycloakID string,
) (*usuario.Usuario, error) {
	row, err := r.queries(ctx).FindByKeycloakID(ctx, keycloakID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, usuario.ErrUserNotFound
		}
		return nil, err
	}

	return mappers.MapFindByKeycloakIDRowToDomain(row), nil
}

func (r *usuarioRepositorySQLC) FindByUsername(
	ctx context.Context,
	username string,
) (*usuario.Usuario, error) {
	row, err := r.queries(ctx).FindByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, usuario.ErrUserNotFound
		}
		return nil, err
	}

	return mappers.MapFindByUsernameRowToDomain(row), nil
}

func (r *usuarioRepositorySQLC) Create(
	ctx context.Context,
	u *usuario.Usuario,
) error {
	if u.Credencial == nil {
		return errors.New("credencial is required")
	}

	q := r.queries(ctx)

	err := q.CreateUsuario(ctx, db.CreateUsuarioParams{
		ID:         u.ID,
		KeycloakID: u.KeycloakID,
		Username:   u.Username,
		Email:      u.Email,
		Nome:       u.Nome,
		NomeFantasia: pgtype.Text{
			String: derefOrEmpty(u.NomeFantasia),
			Valid:  u.NomeFantasia != nil,
		},
	})
	if err != nil {
		return err
	}

	return q.CreateCredencial(ctx, db.CreateCredencialParams{
		ID:              u.Credencial.ID,
		UsuarioID:       u.ID,
		Ativo:           u.Credencial.Ativo,
		PodeAdministrar: u.Credencial.PodeAdministrar,
		PodeVotar:       u.Credencial.PodeVotar,
	})
}

func (r *usuarioRepositorySQLC) UpdateDisplayNamePermissions(
	ctx context.Context,
	userID string,
	displayName *string,
	isActive bool,
	canAdmin bool,
	canVote bool,
) error {
	return r.queries(ctx).UpdateDisplayNamePermissions(ctx, db.UpdateDisplayNamePermissionsParams{
		UserID:      userID,
		DisplayName: derefOrEmpty(displayName),
		IsActive:    isActive,
		CanAdmin:    canAdmin,
		CanVote:     canVote,
	})
}

func (r *usuarioRepositorySQLC) ListUsers(
	ctx context.Context,
	search string,
	page, limit int,
) ([]*usuario.Usuario, int64, error) {
	offset := (page - 1) * limit

	rows, err := r.queries(ctx).ListUsers(ctx, db.ListUsersParams{
		Search:     search,
		LimitRows:  int32(limit),
		OffsetRows: int32(offset),
	})
	if err != nil {
		return nil, 0, fmt.Errorf("ListUsers: %w", err)
	}

	if len(rows) == 0 {
		return []*usuario.Usuario{}, 0, nil
	}

	usuarios := make([]*usuario.Usuario, len(rows))

	for i, row := range rows {
		usuarios[i] = mappers.MapListUsersRowToDomain(row)
	}

	return usuarios, rows[0].TotalCount, nil
}

func (r *usuarioRepositorySQLC) queries(ctx context.Context) *db.Queries {
	if tx, ok := TxFromCtx(ctx); ok {
		return r.q.WithTx(tx)
	}
	return r.q
}

func (r *usuarioRepositorySQLC) FindByID(ctx context.Context, id string) (*usuario.Usuario, error) {
	row, err := r.queries(ctx).FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, usuario.ErrUserNotFound
		}
		return nil, err
	}

	return mappers.MapFindByIDRowToDomain(row), nil
}
