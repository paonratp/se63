import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import WatchVideo from './components/WatchVideo'
import SignIn from './components/SignIn'


import  create_Patientrights from './components/create_patientrights';
import  CreateHistorytaking from './components/HistorytakingCreate';

import  edit from './components/edit';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/watch_video', WatchVideo);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/create_Patientrights', create_Patientrights);
    router.registerRoute('/createHistorytaking', CreateHistorytaking);
  },
});
