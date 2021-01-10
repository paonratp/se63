import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';

import { EntHistorytaking } from '../../api/models/EntHistorytaking';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [Historytakings, setHistorytakings] = useState<EntHistorytaking[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getHistorytakings = async () => {
     const res = await api.listHistorytaking({ offset: 0 });
     setLoading(false);
     setHistorytakings(res);
     console.log(res);
   };
   getHistorytakings();
 }, [loading]);

const deleteHistorytakings = async (id: number) => {
 const res = await api.deleteHistorytaking({ id: id });
 setLoading(true);
};
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">ID</TableCell>
           <TableCell align="center">Hight</TableCell>
           <TableCell align="center">Weight</TableCell>
           <TableCell align="center">T</TableCell>
           <TableCell align="center">P</TableCell>
           <TableCell align="center">R</TableCell>
           <TableCell align="center">BP</TableCell>
           <TableCell align="center">O2</TableCell>
           <TableCell align="center">Symptom</TableCell>
           <TableCell align="center">Nurse</TableCell>
           <TableCell align="center">Symptomseverity</TableCell>
           <TableCell align="center">Department</TableCell>
           <TableCell align="center">Patient</TableCell>
           <TableCell align="center">Datetime</TableCell>
           <TableCell align="center">Manage</TableCell> 
         </TableRow>
       </TableHead>
       <TableBody>
         {Historytakings.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.hight}</TableCell>
             <TableCell align="center">{item.weight}</TableCell>
             <TableCell align="center">{item.temp}</TableCell>
             <TableCell align="center">{item.pulse}</TableCell>
             <TableCell align="center">{item.respiration}</TableCell>
             <TableCell align="center">{item.bp}</TableCell>
             <TableCell align="center">{item.oxygen}</TableCell>
             <TableCell align="center">{item.symptom}</TableCell>
             <TableCell align="center">{item.edges?.nurse?.name}</TableCell>
             <TableCell align="center">{item.edges?.symptomseverity?.symptomseverity}</TableCell>
             <TableCell align="center">{item.edges?.department?.department}</TableCell>
             <TableCell align="center">{item.edges?.patientrecord?.name}</TableCell>
             <TableCell align="center">{item.datetime}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                  deleteHistorytakings(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 DELETE
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
 
