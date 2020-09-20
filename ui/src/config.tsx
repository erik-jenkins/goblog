interface appConfig {
  apiRoot: string;
}

const config: appConfig =
  process.env.NODE_ENV === 'development'
    ? {
        apiRoot: 'https://localhost:4000/',
      }
    : {
        apiRoot: 'https://localhost:3000/',
      };

export default config;
