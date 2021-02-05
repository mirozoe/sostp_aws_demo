import React, { useEffect, useState } from 'react'
import { Typography } from '@material-ui/core'
import { getJournal } from './Network.js'

export const Review = () => {
	const [ items, setItems ] = useState([])

	useEffect( () => {
		setItems( getJournal() )
	}, [])

	console.log(items)
	return (
		<Typography>Účetní denník</Typography>
	)
}
