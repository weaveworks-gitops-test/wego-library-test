import {
  AppContextProvider,
  ApplicationDetail,
  Applications,
  applicationsClient,
  theme,
} from "@weaveworks/weave-gitops";
import * as React from "react";
import {
  BrowserRouter as Router,
  Redirect,
  Route,
  Switch,
} from "react-router-dom";
import { ThemeProvider } from "styled-components";

const APPS_ROUTE = "/my_custom_applications_route";
const APP_DETAIL_ROTUE = "/my_special_application_detail_route";

// Use this to reconcile differences in routing so that links will work correctly.
const linkResolver = (incoming: string): string => {
  const parsed = new URL(incoming, window.location.href);

  switch (parsed.pathname) {
    case "/applications":
      return `${APPS_ROUTE}${parsed.search}`;

    case "/application_detail":
      return `${APP_DETAIL_ROTUE}${parsed.search}`;

    default:
      return incoming;
  }
};

export default function App() {
  return (
    <div>
      My app
      <ThemeProvider theme={theme}>
        <h3>My custom App!!</h3>
        <div>
          <Router>
            <AppContextProvider
              linkResolver={linkResolver}
              applicationsClient={applicationsClient}
            >
              <Switch>
                <Route exact path={APPS_ROUTE} component={Applications} />
                <Route
                  exact
                  path={APP_DETAIL_ROTUE}
                  component={ApplicationDetail}
                />
                <Redirect from="/" to={APPS_ROUTE} />
                <Route exact path="*" component={() => <h3>404</h3>} />
              </Switch>
            </AppContextProvider>
          </Router>
        </div>
      </ThemeProvider>
    </div>
  );
}
