import logo from './logo.svg';
import React, { useState } from 'react'
import { Box, Button, Grid, Paper, Typography} from '@material-ui/core'
import { InputScreen } from './InputScreen.js'
import { Review } from './Review.js'
import './App.css';


function App() {
 const [ screen, setScreen ] = useState(null)

	const selectScreen = () => {
		switch (screen) {
			case "InputScreen":
				return ( <InputScreen /> )
				break
			case "Review":
				return ( <Review /> )
				break
			default:
				return ( <Paper style={{ 
												backgroundImage: "url(/nick-morrison-FHnnjk1Yj7Y-unsplash.jpg)", 
												backgroundSize: "100%", 
												width: "724px", 
												height: "450px",
												flexDirection: "column",
												display: "flex",				
								}} >
									<Typography variant='h2' >Účetnictví Tomáš Marný</Typography>
								</Paper> )
		}
	}

  return (
    <div className="App">
      <header className="App-header" >
					<Grid container justify="center" spacing={2} >
						<Grid item sm={12} >
							<Box mt={3}>
							<Grid container justify="center" spacing={2} >
								<Grid item	>
									<Button 
										variant="contained" 
										color="secondary"
										onClick={() => setScreen("Review")}	
									>Účetní kniha</Button>
								</Grid>
								<Grid item	>
									<Button 
										variant="contained" 
										color="secondary" 
										onClick={() => setScreen("InputScreen")}
									>Vkládání účtů</Button>
								</Grid>
							</Grid>
							</Box>
						</Grid>
						<Grid item>
							<Box m={2} >
								{selectScreen()}				
							</Box>
						</Grid>
					</Grid>
      </header>
    </div>
  );
}

export default App;
