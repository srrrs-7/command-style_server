package graph

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sRRRs-7/loose_style.git/cfg"
	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/dataloaders"
	"github.com/sRRRs-7/loose_style.git/graph/generated"
	"github.com/sRRRs-7/loose_style.git/token"
)

// create resolver instance
func NewResolver(config cfg.Config, store db.Store, pool *pgxpool.Pool) (*Resolver, token.Maker, error) {
	// tokenMaker instance
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	// initialize dataloaders instance
	dl := dataloaders.NewRetriever()
	// api instance
	resolver := &Resolver{
		store:       store,
		tokenMaker:  tokenMaker,
		config:      config,
		dataloaders: *dl,
		tx:          pool,
	}

	return resolver, tokenMaker, nil
}

// Define the Graphql handler
func graphqlHandler(r *Resolver) gin.HandlerFunc {
	// context store instance
	h := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &Resolver{
				store:       r.store,
				tokenMaker:  r.tokenMaker,
				config:      r.config,
				dataloaders: r.dataloaders,
				tx:          r.tx,
			}}))
	h.AddTransport(transport.POST{})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/admin/query") // playground fetcher endpoint

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
