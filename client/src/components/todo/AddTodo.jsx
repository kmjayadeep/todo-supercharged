import Card from "@material-ui/core/Card";
import CardActions from "@material-ui/core/CardActions";
import CardContent from "@material-ui/core/CardContent";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import TextField from "@material-ui/core/TextField";
import SaveIcon from "@material-ui/icons/Save";
// import { addTodo } from "./todo.svc";

import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles({
  root: {
    minWidth: 275,
    margin: "1rem",
  },
  bullet: {
    display: "inline-block",
    margin: "0 2px",
    transform: "scale(0.8)",
  },
  title: {
    fontSize: 14,
    marginBottom: 10,
  },
  pos: {
    marginBottom: 12,
  },
  done: {
    color: "#0db70d",
  },
  add: {
    fontSize: "5rem",
  },
  inputtitle: {
    marginBottom: 10,
  },
});

export function AddTodo({ refresh }) {
  const classes = useStyles();

  return (
    <Card className={classes.root} color="secondary">
      <CardContent>
        <Typography variant="h5" component="h2" className={classes.title}>
          <b>New Todo</b>
        </Typography>
        <TextField id="title" label="Title" className={classes.inputtitle} fullWidth/>
        <TextField id="description" label="Description" multiline rows="4" fullWidth/>
      </CardContent>
      <CardActions>
        <Button size="small" color="primary" onClick={() => {}}>
          <SaveIcon />
          Save
        </Button>
      </CardActions>
    </Card>
  );
}
