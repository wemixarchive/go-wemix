// rocksdb-c-util.go
// +build rocksdb

package rocksdb

/*

#include <stdlib.h>
#include "rocksdb/c.h"

extern void replayPut(char *, char *, size_t, char *, size_t);
extern void replayDel(char *, char *, size_t);

static void replay_put(void *ptr, const char *k, size_t klen, const char *v, size_t vlen)
{
    replayPut((char *) ptr, (char *) k, klen, (char *) v, vlen);
}

static void replay_del(void *ptr, const char *k, size_t klen)
{
    replayDel((char *) ptr, (char *) k, klen);
}

void replay_iterate(int *wp, void *bp)
{
    rocksdb_writebatch_iterate(bp, wp, replay_put, replay_del);
}

*/
import "C"
