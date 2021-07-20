import Card from "@material-ui/core/Card";
import CardActions from "@material-ui/core/CardActions";
import CardContent from "@material-ui/core/CardContent";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import DoneIcon from "@material-ui/icons/Done";
import DeleteIcon from "@material-ui/icons/Delete";
import { markAsDone, deleteTodo } from "./todo.svc";

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
  },
  pos: {
    marginBottom: 12,
  },
  done: {
    color: "#0db70d"
  }
});

export function Todo({ todo, refresh }) {
  const classes = useStyles();

  let cls = ""
  if(todo.completed) {
    cls = cls + " " + classes.done
  }

  return (
    <Card className={classes.root} color="secondary">
      <CardContent>
        <Typography variant="h5" component="h2" className={cls}>
    <b>
          {todo.title}</b>
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          {new Date(todo.createdAt).toLocaleString()}
        </Typography>
        <Typography variant="body2" component="p">
          {todo.description}
        </Typography>
      </CardContent>
      <CardActions>
        {todo.completed || (
          <Button
            size="small"
            color="primary"
            onClick={() => {
              markAsDone(todo._id).then(() => refresh());
            }}
          >
            <DoneIcon />
            Mark as done
          </Button>
        )}
        <Button
          size="small"
          color="secondary"
          onClick={() => {
            deleteTodo(todo._id).then(() => refresh());
          }}
        >
          <DeleteIcon />
          Delete
        </Button>
      </CardActions>
    </Card>
  );
}
