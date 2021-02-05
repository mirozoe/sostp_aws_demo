import React, { useEffect, useState } from 'react'
import { Paper, Table, TableContainer, TableHead, TableBody, TableRow, TableCell, Typography } from '@material-ui/core'
import { getJournal } from './Network.js'

export const Review = () => {
	const [ items, setItems ] = useState([])

	useEffect(async () => {
		setItems(await getJournal())
	}, [])

	items.map( i => console.log(i))
	return (
		<>
		<Typography>Účetní denník</Typography>
		<TableContainer component={Paper}>
			<Table size="small">
				<TableHead>
					<TableRow>
						<TableCell>Datum</TableCell>
						<TableCell>Popis</TableCell>
						<TableCell>Typ</TableCell>
						<TableCell>Cena</TableCell>
						<TableCell>MD</TableCell>
						<TableCell>D</TableCell>
					</TableRow>
				</TableHead>
				<TableBody>
				{ items.map( i => (
					<TableRow key={i.itemid}>
						<TableCell>{i.date}</TableCell>
						<TableCell>{i.description}</TableCell>
						<TableCell>{i.type}</TableCell>
						<TableCell>{i.price}</TableCell>
						<TableCell>{i.debit}</TableCell>
						<TableCell>{i.kredit}</TableCell>
					</TableRow>
				))}
				</TableBody>
			</Table>
		</TableContainer>
		</>
	)
}
