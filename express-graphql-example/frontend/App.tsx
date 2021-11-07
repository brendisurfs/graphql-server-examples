import { gql, useQuery } from "@apollo/client";
import React from "react";

const query = gql`
	{
		hello
	}
`;

const App = () => {
	const { loading, error, data } = useQuery(query);
	if (loading) return <>Loading</>;
	if (error) return <>{JSON.stringify(error)}</>;

	let objData = JSON.stringify(data);

	return <div>{data.hello}</div>;
};

export default App;
