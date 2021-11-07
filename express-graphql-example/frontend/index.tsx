import ReactDOM from "react-dom";
import {
	ApolloCache,
	ApolloClient,
	ApolloProvider,
	InMemoryCache,
} from "@apollo/client";
import App from "./App";

const client = new ApolloClient({
	uri: "http://localhost:4000/graphql",
	cache: new InMemoryCache(),
});

ReactDOM.render(
	<ApolloProvider client={client}>
		<App />
	</ApolloProvider>,
	document.getElementById("root")
);
