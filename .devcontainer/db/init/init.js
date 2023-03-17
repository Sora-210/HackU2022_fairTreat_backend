rs.initiate({
    _id: 'rs0',
    members: [
    {_id: 0, host: "db-primary:27017"},
    {_id: 1, host: "db-secondary:27017"},
    {_id: 2, host: "db-arbiter:27017"}
    ],
});  