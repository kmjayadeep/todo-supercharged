import Card from "@material-ui/core/Card";
import CardActions from "@material-ui/core/CardActions";
import CardContent from "@material-ui/core/CardContent";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import DoneIcon from "@material-ui/icons/Done";
import DeleteIcon from "@material-ui/icons/Delete";
import { markAsDone } from "./todo.svc";

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
});

export function Todo({ todo, refresh }) {
  const classes = useStyles();

  return (
    <Card className={classes.root} color="secondary">
      <CardContent>
        <Typography variant="h5" component="h2">
          {todo.title}
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
          <Button size="small" color="primary" onClick={()=>{
            markAsDone(todo._id).then(()=>refresh());
          }}>
            <DoneIcon />
            Mark as done
          </Button>
        )}
        <Button size="small" color="secondary">
          <DeleteIcon />
          Delete
        </Button>
      </CardActions>
    </Card>
  );
}
