import express from "express";
import { graphqlHTTP } from "express-graphql";
import { buildSchema } from "graphql";
import cors from "cors";
// make a schema
let schema = buildSchema(
	`
    type Query {
        hello: String
    }
    `
);

let root = {
	hello: () => {
		return "Hello Nerds";
	},
};

let app = express();
app.use(cors());
app.use(
	"/graphql",
	graphqlHTTP({
		schema,
		rootValue: root,
	})
);

app.listen(4000);
console.log("running gql on http://localhost:4000/graphql");
