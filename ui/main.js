import React from "react";
import ReactDOM from "react-dom";
import MyApp from "./MyApp.tsx";

// eslint-disable-next-line
ReactDOM.render(<MyApp />, document.getElementById("app"));
// eslint-disable-next-line
if (module.hot) {
  // eslint-disable-next-line
  module.hot.accept();
}
