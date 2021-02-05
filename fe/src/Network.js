const baseURL = "https://5u0ws0u3fb.execute-api.us-east-2.amazonaws.com/prod"

export const getJournal = async () => {
	const result = await fetch( `${baseURL}/journal`, {
		method: 'GET',
	})
	console.log(result)
	if (result.status === 200) {
		return await result.json()
	}
	return []
}
