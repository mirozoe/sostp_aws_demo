import React, { useState } from 'react'
import { Button, MenuItem, Paper, Select, Snackbar, Table, TableBody, TableCell, TableContainer, TableRow, TextField, Typography } from '@material-ui/core'
import { insertJournal } from './Network.js'


export const InputScreen = () => {
	const [ date, setDate ] = useState(null)
	const [ desc, setDesc ] = useState(null)
	const [ type, setType ] = useState("FP")
	const [ price, setPrice ] = useState(null)
	const [ debit, setDebit ] = useState(null)
	const [ kredit, setKredit ] = useState(null)
	const [ dateErr, setDateErr ] = useState(false)
	const [ descErr, setDescErr ] = useState(false)
	const [ priceErr, setPriceErr ] = useState(false)
	const [ debitErr, setDebitErr ] = useState(false)
	const [ kreditErr, setKreditErr ] = useState(false)
	const [ saved, setSaved ] = useState(null)


	const zeroizeErrors = () => {
		setDateErr(false)
		setDescErr(false)
		setPriceErr(false)
		setDebitErr(false)
		setKreditErr(false)
	}
	
	const zeroizeValues = () => {
		setDate(null)
		setDesc(null)
		setType("FP")
		setPrice(null)
		setDebit(null)
		setKredit(null)
	}

	const validate = () => {
		zeroizeErrors()
		if (date === null) { 
			setDateErr(true)
			return false
		}
		if (desc === null) { 
			setDescErr(true)
			return false
		}
		if (price === null) { 
			setPriceErr(true)
			return false
		}
		if (debit === null) { 
			setDebitErr(true)
			return false
		}
		if (kredit === null) { 
			setKreditErr(true)
			return false
		}
		return true
	}

	const save = async () => {
		if (validate()) {
			console.log(`${date}-${desc}-${type}-${price}-${debit}-${kredit}`)
			const d = {
				date: date,
				description: desc,
				type: type,
				price: price,
				debit: debit,
				kredit: kredit
			}
			const result = await insertJournal(d)
			setSaved(result)
			zeroizeErrors()
			zeroizeValues()
		}
	}

	
	const handleType = event => setType(event.target.value)

	return (
		<>
		<Typography>Vkládání účetních dokladů</Typography>
		<TableContainer component={Paper}>
			<Table>
				<TableBody>
					<TableRow>
						<TableCell>Datum</TableCell>
						<TableCell>
							<TextField 
								onChange={ (d) => setDate(d.target.value) }
								error={ dateErr }
								helperText={ dateErr ? "Zadejte datum dokladu" : null }
							/>
						</TableCell>
					</TableRow>
					<TableRow>
						<TableCell>Popis</TableCell>
						<TableCell>
							<TextField 
								onChange={ (d) => setDesc(d.target.value) }
								error={ descErr }				
								helperText={ descErr ? "Zadejte popis účetní operace" : null }
							/>
						</TableCell>
					</TableRow>
					<TableRow>
						<TableCell>Typ</TableCell>
						<TableCell>
							<Select
								value={ type }
								onChange={ handleType }
							>
								<MenuItem value={"FP"}>FP</MenuItem>
								<MenuItem value={"FV"}>FV</MenuItem>
								<MenuItem value={"PDP"}>PDP</MenuItem>
								<MenuItem value={"PDV"}>PDV</MenuItem>
								<MenuItem value={"VDV"}>VDV</MenuItem>
								<MenuItem value={"VDP"}>VDP</MenuItem>
							</Select>
						</TableCell>
					</TableRow>
					<TableRow>
						<TableCell>Cena</TableCell>
						<TableCell>
							<TextField 
								onChange={ (d) => setPrice(d.target.value) }
								error={ priceErr }
								helperText={ priceErr ? "Zadejte částku operace" : null }
							/>
						</TableCell>
					</TableRow>
					<TableRow>
						<TableCell>Má dáti</TableCell>
						<TableCell>
							<TextField 
								onChange={ (d) => setDebit(d.target.value) }
								error={ debitErr }
								helperText={ debitErr ? "Zadajte účet stranu má dáti" : null }
							/>
						</TableCell>
					</TableRow>
					<TableRow>
						<TableCell>Dal</TableCell>
						<TableCell>
							<TextField 
								onChange={ (d) => setKredit(d.target.value) }
								error={ kreditErr }
								helperText={ kreditErr ? "Zadejte účet stranu dal" : null }
							/>
						</TableCell>
					</TableRow>
					<TableRow>
						<TableCell 
							colSpan={2}
							align="center"
						>
							<Button onClick={ () => save() }>
								Odeslat
							</Button>
						</TableCell>
					</TableRow>
				</TableBody>
			</Table>
		</TableContainer>
		<Snackbar open={saved} autoHideDuration={6000} onClose={() => setSaved(false)}>
        Uloženo
    </Snackbar>
		</>
	)
}
