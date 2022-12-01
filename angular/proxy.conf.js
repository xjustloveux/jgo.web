const PROXY_CONFIG = [
  {
    context: [
      "/shared", "/a", "/b", "/c"
    ],
    target: "http://localhost:8080",
    secure: false,
    changeOrigin: true
  }
];

module.exports = PROXY_CONFIG;
