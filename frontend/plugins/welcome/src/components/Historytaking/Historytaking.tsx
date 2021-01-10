import React from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../tableHistorytaking';
import Button from '@material-ui/core/Button';
import { Typography, Grid } from '@material-ui/core';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
export default function HistorytakingTable() {
const profile = { givenName: '' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
      title={`${profile.givenName || 'HISTORYTAKING DEPARTMENT'}`}
      subtitle=""
     ></Header>
     <Content>
       <ContentHeader title="">
       <Grid item  xs={12} md={12}>
       <Typography variant="h6" gutterBottom align="center">
         <div>
          <Link component={RouterLink} to="/createHistorytaking">
            <Button variant="contained" color="primary" style={{backgroundColor: "#21b6ae"}}>
              ADD DATA
            </Button>
          </Link>
         </div>
      </Typography>
        </Grid>
       </ContentHeader>
       <div>
       </div>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};

