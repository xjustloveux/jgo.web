const PROXY_CONFIG = [
  {
    context: [
      "/Shared", "/A", "/B", "/C"
    ],
    target: "http://localhost:8080",
    secure: false,
    changeOrigin: true
  }
];

module.exports = PROXY_CONFIG;
