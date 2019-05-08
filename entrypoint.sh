#!/bin/sh
set -e

# first arg is `-f` or `--some-option`
# or first arg is `something.toml`
if [ "${1#-}" != "$1" ] || [ "${1%.toml}" != "$1" ]; then
	set -- /bin/redis-failover "$@"
fi

# allow the container to be started with `--user`
if [ "$1" = 'redis-failover' -a "$(id -u)" = '0' ]; then
	chown -R ledis /datastore
    chown ledis:ledis /bin/redis-*
	exec gosu ledis "$0" "$@"
fi

exec "$@"