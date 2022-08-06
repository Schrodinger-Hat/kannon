// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package sqlc

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createDomainStmt, err = db.PrepareContext(ctx, createDomain); err != nil {
		return nil, fmt.Errorf("error preparing query CreateDomain: %w", err)
	}
	if q.createMessageStmt, err = db.PrepareContext(ctx, createMessage); err != nil {
		return nil, fmt.Errorf("error preparing query CreateMessage: %w", err)
	}
	if q.createPoolStmt, err = db.PrepareContext(ctx, createPool); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePool: %w", err)
	}
	if q.createPoolWithFieldsStmt, err = db.PrepareContext(ctx, createPoolWithFields); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePoolWithFields: %w", err)
	}
	if q.createStatsKeysStmt, err = db.PrepareContext(ctx, createStatsKeys); err != nil {
		return nil, fmt.Errorf("error preparing query CreateStatsKeys: %w", err)
	}
	if q.createTemplateStmt, err = db.PrepareContext(ctx, createTemplate); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTemplate: %w", err)
	}
	if q.findDomainStmt, err = db.PrepareContext(ctx, findDomain); err != nil {
		return nil, fmt.Errorf("error preparing query FindDomain: %w", err)
	}
	if q.findDomainWithKeyStmt, err = db.PrepareContext(ctx, findDomainWithKey); err != nil {
		return nil, fmt.Errorf("error preparing query FindDomainWithKey: %w", err)
	}
	if q.findTemplateStmt, err = db.PrepareContext(ctx, findTemplate); err != nil {
		return nil, fmt.Errorf("error preparing query FindTemplate: %w", err)
	}
	if q.getAllDomainsStmt, err = db.PrepareContext(ctx, getAllDomains); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllDomains: %w", err)
	}
	if q.getDomainsStmt, err = db.PrepareContext(ctx, getDomains); err != nil {
		return nil, fmt.Errorf("error preparing query GetDomains: %w", err)
	}
	if q.getSendingDataStmt, err = db.PrepareContext(ctx, getSendingData); err != nil {
		return nil, fmt.Errorf("error preparing query GetSendingData: %w", err)
	}
	if q.getValidPublicStatsKeyByKidStmt, err = db.PrepareContext(ctx, getValidPublicStatsKeyByKid); err != nil {
		return nil, fmt.Errorf("error preparing query GetValidPublicStatsKeyByKid: %w", err)
	}
	if q.getValidStatsKeysStmt, err = db.PrepareContext(ctx, getValidStatsKeys); err != nil {
		return nil, fmt.Errorf("error preparing query GetValidStatsKeys: %w", err)
	}
	if q.prepareForSendStmt, err = db.PrepareContext(ctx, prepareForSend); err != nil {
		return nil, fmt.Errorf("error preparing query PrepareForSend: %w", err)
	}
	if q.setDomainKeyStmt, err = db.PrepareContext(ctx, setDomainKey); err != nil {
		return nil, fmt.Errorf("error preparing query SetDomainKey: %w", err)
	}
	if q.setSendingPoolDeliveredStmt, err = db.PrepareContext(ctx, setSendingPoolDelivered); err != nil {
		return nil, fmt.Errorf("error preparing query SetSendingPoolDelivered: %w", err)
	}
	if q.setSendingPoolErrorStmt, err = db.PrepareContext(ctx, setSendingPoolError); err != nil {
		return nil, fmt.Errorf("error preparing query SetSendingPoolError: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createDomainStmt != nil {
		if cerr := q.createDomainStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createDomainStmt: %w", cerr)
		}
	}
	if q.createMessageStmt != nil {
		if cerr := q.createMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createMessageStmt: %w", cerr)
		}
	}
	if q.createPoolStmt != nil {
		if cerr := q.createPoolStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPoolStmt: %w", cerr)
		}
	}
	if q.createPoolWithFieldsStmt != nil {
		if cerr := q.createPoolWithFieldsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPoolWithFieldsStmt: %w", cerr)
		}
	}
	if q.createStatsKeysStmt != nil {
		if cerr := q.createStatsKeysStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createStatsKeysStmt: %w", cerr)
		}
	}
	if q.createTemplateStmt != nil {
		if cerr := q.createTemplateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTemplateStmt: %w", cerr)
		}
	}
	if q.findDomainStmt != nil {
		if cerr := q.findDomainStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findDomainStmt: %w", cerr)
		}
	}
	if q.findDomainWithKeyStmt != nil {
		if cerr := q.findDomainWithKeyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findDomainWithKeyStmt: %w", cerr)
		}
	}
	if q.findTemplateStmt != nil {
		if cerr := q.findTemplateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findTemplateStmt: %w", cerr)
		}
	}
	if q.getAllDomainsStmt != nil {
		if cerr := q.getAllDomainsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllDomainsStmt: %w", cerr)
		}
	}
	if q.getDomainsStmt != nil {
		if cerr := q.getDomainsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDomainsStmt: %w", cerr)
		}
	}
	if q.getSendingDataStmt != nil {
		if cerr := q.getSendingDataStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSendingDataStmt: %w", cerr)
		}
	}
	if q.getValidPublicStatsKeyByKidStmt != nil {
		if cerr := q.getValidPublicStatsKeyByKidStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getValidPublicStatsKeyByKidStmt: %w", cerr)
		}
	}
	if q.getValidStatsKeysStmt != nil {
		if cerr := q.getValidStatsKeysStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getValidStatsKeysStmt: %w", cerr)
		}
	}
	if q.prepareForSendStmt != nil {
		if cerr := q.prepareForSendStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing prepareForSendStmt: %w", cerr)
		}
	}
	if q.setDomainKeyStmt != nil {
		if cerr := q.setDomainKeyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setDomainKeyStmt: %w", cerr)
		}
	}
	if q.setSendingPoolDeliveredStmt != nil {
		if cerr := q.setSendingPoolDeliveredStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setSendingPoolDeliveredStmt: %w", cerr)
		}
	}
	if q.setSendingPoolErrorStmt != nil {
		if cerr := q.setSendingPoolErrorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setSendingPoolErrorStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                              DBTX
	tx                              *sql.Tx
	createDomainStmt                *sql.Stmt
	createMessageStmt               *sql.Stmt
	createPoolStmt                  *sql.Stmt
	createPoolWithFieldsStmt        *sql.Stmt
	createStatsKeysStmt             *sql.Stmt
	createTemplateStmt              *sql.Stmt
	findDomainStmt                  *sql.Stmt
	findDomainWithKeyStmt           *sql.Stmt
	findTemplateStmt                *sql.Stmt
	getAllDomainsStmt               *sql.Stmt
	getDomainsStmt                  *sql.Stmt
	getSendingDataStmt              *sql.Stmt
	getValidPublicStatsKeyByKidStmt *sql.Stmt
	getValidStatsKeysStmt           *sql.Stmt
	prepareForSendStmt              *sql.Stmt
	setDomainKeyStmt                *sql.Stmt
	setSendingPoolDeliveredStmt     *sql.Stmt
	setSendingPoolErrorStmt         *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                              tx,
		tx:                              tx,
		createDomainStmt:                q.createDomainStmt,
		createMessageStmt:               q.createMessageStmt,
		createPoolStmt:                  q.createPoolStmt,
		createPoolWithFieldsStmt:        q.createPoolWithFieldsStmt,
		createStatsKeysStmt:             q.createStatsKeysStmt,
		createTemplateStmt:              q.createTemplateStmt,
		findDomainStmt:                  q.findDomainStmt,
		findDomainWithKeyStmt:           q.findDomainWithKeyStmt,
		findTemplateStmt:                q.findTemplateStmt,
		getAllDomainsStmt:               q.getAllDomainsStmt,
		getDomainsStmt:                  q.getDomainsStmt,
		getSendingDataStmt:              q.getSendingDataStmt,
		getValidPublicStatsKeyByKidStmt: q.getValidPublicStatsKeyByKidStmt,
		getValidStatsKeysStmt:           q.getValidStatsKeysStmt,
		prepareForSendStmt:              q.prepareForSendStmt,
		setDomainKeyStmt:                q.setDomainKeyStmt,
		setSendingPoolDeliveredStmt:     q.setSendingPoolDeliveredStmt,
		setSendingPoolErrorStmt:         q.setSendingPoolErrorStmt,
	}
}
