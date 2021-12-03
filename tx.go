package rds

import "github.com/aws/aws-sdk-go/service/rdsdataservice"

type tx struct {
	conn   *conn
	output *rdsdataservice.BeginTransactionOutput
}

func (tx *tx) Commit() error {
	_, err := tx.conn.rds.CommitTransaction(&rdsdataservice.CommitTransactionInput{
		ResourceArn:   &tx.conn.resourceArn,
		SecretArn:     &tx.conn.secretArn,
		TransactionId: tx.output.TransactionId,
	})
	return err
}

func (tx *tx) Rollback() error {
	_, err := tx.conn.rds.RollbackTransaction(&rdsdataservice.RollbackTransactionInput{
		ResourceArn:   &tx.conn.resourceArn,
		SecretArn:     &tx.conn.secretArn,
		TransactionId: tx.output.TransactionId,
	})
	return err
}
