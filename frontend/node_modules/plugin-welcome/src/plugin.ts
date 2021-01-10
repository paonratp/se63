import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import WatchVideo from './components/WatchVideo'
import SignIn from './components/SignIn'


import  create_Patientrights from './components/create_patientrights';
import  createHistorytaking from './components/edit';
import  Historytaking from './components/Historytaking';
import  tableHistorytaking from './components/tableHistorytaking';

import  edit from './components/Login';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Historytaking);
    router.registerRoute('/watch_video', WatchVideo);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/create_Patientrights', create_Patientrights);
    router.registerRoute('/createHistorytaking', createHistorytaking);
  },
});
