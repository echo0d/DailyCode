import React from 'react';
import ReactDOM from 'react-dom/client';
import 'antd/dist/reset.css';
import AppLayout from './layout/AppLayout';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <AppLayout />
  </React.StrictMode>
);
