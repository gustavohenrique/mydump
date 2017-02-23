# mydump

> A simple way to export your MySQL database

## Usage

1. Create the `config.json` file like bellow:

```
{
    "host": "192.168.99.100",
    "port": "3306",
    "database": "mydatabase",
    "user": "root",
    "password": "root"
}
```

2. Run the `mydump` binary according of your OS:

```
./mydump-linux-amd64
```

3. The dump file generated will be placed in the `dumps` directory.

## License

MIT

