import { Container } from "@material-ui/core";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import { TodoList } from "./components/todo/TodoList";

import "./App.css";

function App() {
  return (
    <Container maxWidth="lg">
      <AppBar position="static">
        <Toolbar>
          <Typography variant="subtitle1">Todo</Typography>
          <div>&nbsp;</div>
          <Typography variant="h6">Supercharged</Typography>
        </Toolbar>
      </AppBar>
      <TodoList/>
    </Container>
  );
}

export default App;
