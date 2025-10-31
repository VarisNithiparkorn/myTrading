import localFont from 'next/font/local';

export const myLocalFont = localFont({
  src: [
    {
      path: '../../public/font/Santhai/IBMPlexSansThai-Medium.ttf',
    weight: '500', 
      style: 'normal',
    },

  ],
  variable: '--font-mycustom',
  display: 'swap',
});