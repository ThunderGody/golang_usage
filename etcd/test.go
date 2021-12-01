package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"strconv"
	"time"
)

var (
	dialTimeout = 2 * time.Second
	requestTimeout = 10 * time.Second
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
	cli, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints: []string{"47.104.177.252:2379"},
	})
	defer cli.Close()
	kv := clientv3.NewKV(cli)

	GetSingleValueDemo(ctx, kv)
	GetMultipleValuesWithPaginationDemo(ctx, kv)
	LeaseDemo(ctx, cli, kv)
}
func GetSingleValueDemo(ctx context.Context, kv clientv3.KV) {
	fmt.Println("*** GetSingleValueDemo ***")
	kv.Delete(ctx, "key", clientv3.WithPrefix())

	pr, _ := kv.Put(ctx, "key", "444")
	rev := pr.Header.Revision
	fmt.Println("Revision: ", rev)

	gr, _ := kv.Get(ctx, "key")
	fmt.Println("Value: ", string(gr.Kvs[0].Value), "Revision: ", gr.Header.Revision)

	kv.Put(ctx, "key", "555")

	gr, _ = kv.Get(ctx, "key")
	fmt.Println("Value: ", string(gr.Kvs[0].Value), "Revision: ", gr.Header.Revision)

}

func GetMultipleValuesWithPaginationDemo(ctx context.Context, kv clientv3.KV) {
	fmt.Println("*** GetMultipleValuesWithPaginationDemo ***")
	kv.Delete(ctx, "key", clientv3.WithPrefix())

	for i := 0; i < 10; i++ {
		k := fmt.Sprintf("key_%02d", i)
		kv.Put(ctx, k, strconv.Itoa(i))
	}

	opts := []clientv3.OpOption {
		clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend),
		clientv3.WithLimit(3),
	}

	gr, _ := kv.Get(ctx, "key", opts...)
	fmt.Println("--- First page ---")
	for _, item := range gr.Kvs {
		fmt.Println(string(item.Key), string(item.Value))
	}

	lastKey := string(gr.Kvs[len(gr.Kvs)-1].Key)
	fmt.Println("---------- Second page ----")
	opts = append(opts, clientv3.WithFromKey())

	gr, _ = kv.Get(ctx, lastKey, opts...)
	for _, item := range gr.Kvs {
		fmt.Println(string(item.Key), string(item.Value))
	}
}

func LeaseDemo(ctx context.Context, cli *clientv3.Client, kv clientv3.KV) {
	fmt.Println("*** LeaseDemo ***")
	kv.Delete(ctx, "key", clientv3.WithPrefix())

	gr, _ := kv.Get(ctx, "key")
	if len(gr.Kvs) == 0 {
		fmt.Println("No 'key'")
	}

	lease, _ := cli.Grant(ctx, 1)
	kv.Put(ctx, "key", "value", clientv3.WithLease(lease.ID))

	gr, _ = kv.Get(ctx, "key")
	if len(gr.Kvs) == 1 {
		fmt.Println("Found 'key'")
	}

	// Let the TTL expire
	time.Sleep(3 * time.Second)

	gr, _ = kv.Get(ctx, "key")
	if len(gr.Kvs) == 0 {
		fmt.Println("No more 'key'")
	}
}