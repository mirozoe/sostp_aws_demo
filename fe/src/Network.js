const baseURL = "https://5u0ws0u3fb.execute-api.us-east-2.amazonaws.com/prod"

export const getJournal = async (token) => {
	const result = await fetch( `${baseURL}/journal`, {
		method: 'POST',
		headers: {
			"Authorization": `Bearer ${token}`
		}
	})
	if (result.status === 200) {
		const j = await result.json()
		return j
	}
	return []
}

export const insertJournal = async (token, data) => {
	console.log(data)
	const result = await fetch( `${baseURL}/insert`, {
		method: 'POST',
		headers: {
			'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json'
		},
		body: JSON.stringify({
			description: data.description,
			date: data.date,
			type: data.type,
			price: parseInt(data.price, 10),
			debit: parseInt(data.debit, 10),
			kredit: parseInt(data.kredit, 10)
		}),
	})
	if (result.status === 200) {
		return true
	}
	return false
}
