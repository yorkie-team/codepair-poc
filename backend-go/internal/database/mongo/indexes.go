package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Collection names based on your schema.
const (
	ColUsers                     = "users"
	ColUserWorkspaces            = "user_workspaces"
	ColWorkspaces                = "workspaces"
	ColDocuments                 = "documents"
	ColWorkspaceInvitationTokens = "workspace_invitation_tokens"
	ColDocumentSharingTokens     = "document_sharing_tokens"
	ColIntelligenceLogs          = "intelligence_logs"
)

// collectionInfo holds the name of the collection and indexes to create.
type collectionInfo struct {
	name    string
	indexes []mongo.IndexModel
}

// collectionInfos holds all the indexes you want to create for each collection.
var collectionInfos = []collectionInfo{
	{
		name:    ColUsers,
		indexes: nil, // Default _id index used
	},
	{
		name:    ColUserWorkspaces,
		indexes: nil, // Default _id index used
	},
	{
		name:    ColWorkspaces,
		indexes: nil, // Default _id index used
	},
	{
		name:    ColDocuments,
		indexes: nil, // Default _id index used
	},
	{
		name:    ColWorkspaceInvitationTokens,
		indexes: nil, // Default _id index used
	},
	{
		name:    ColDocumentSharingTokens,
		indexes: nil, // Default _id index used
	},
	{
		name:    ColIntelligenceLogs,
		indexes: nil, // Default _id index used
	},
}

// ensureIndexes iterates over collectionInfos and creates the indexes defined.
func ensureIndexes(ctx context.Context, db *mongo.Database) error {
	for _, info := range collectionInfos {
		_, err := db.Collection(info.name).Indexes().CreateMany(ctx, info.indexes)
		if err != nil {
			return fmt.Errorf("failed to create indexes on %s: %w", info.name, err)
		}
	}
	return nil
}
