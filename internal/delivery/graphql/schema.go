package graphql

import (
	"github.com/sorroche-m/desafio-clean-architecture/internal/domain"

	"github.com/graphql-go/graphql"
)

func NewSchema(orderUseCase domain.OrderUseCase) (graphql.Schema, error) {
	orderType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"customer_id": &graphql.Field{
				Type: graphql.String,
			},
			"amount": &graphql.Field{
				Type: graphql.Float,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
			"updated_at": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"orders": &graphql.Field{
				Type: graphql.NewList(orderType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return orderUseCase.ListOrders()
				},
			},
			"order": &graphql.Field{
				Type: orderType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					return orderUseCase.GetOrder(id)
				},
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createOrder": &graphql.Field{
				Type: orderType,
				Args: graphql.FieldConfigArgument{
					"customer_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"amount": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					customerID := p.Args["customer_id"].(string)
					amount := p.Args["amount"].(float64)
					return orderUseCase.CreateOrder(customerID, amount)
				},
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}
