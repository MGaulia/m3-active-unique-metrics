	"github.com/m3db/m3/src/dbnode/encoding/proto"
	xconfig "github.com/m3db/m3/src/x/config"
	"github.com/m3db/m3/src/x/context"
	"github.com/m3db/m3/src/x/ident"
	"github.com/m3db/m3/src/x/instrument"
	xlog "github.com/m3db/m3/src/x/log"
	"github.com/m3db/m3/src/x/pool"
	xsync "github.com/m3db/m3/src/x/sync"
	"github.com/jhump/protoreflect/desc"
	var schema *desc.MessageDescriptor
	if cfg.Proto != nil {
		logger.Info("Probuf data mode enabled")
		schema, err = proto.ParseProtoSchema(cfg.Proto.SchemaFilePath)
		if err != nil {
			logger.Fatalf("error parsing protobuffer schema: %v", err)
		}
	}

	opts = withEncodingAndPoolingOptions(cfg, logger, schema, opts, cfg.PoolingPolicy)
		},
		func(opts client.AdminOptions) client.AdminOptions {
			if cfg.Proto != nil {
				return opts.SetEncodingProto(
					schema,
					encoding.NewOptions(),
				).(client.AdminOptions)
			}
			return opts
		},
	)
	schema *desc.MessageDescriptor,
		if schema != nil {
			enc := proto.NewEncoder(time.Time{}, encodingOpts)
			enc.SetSchema(schema)
			return enc
		}

		if schema != nil {
			return proto.NewIterator(r, schema, encodingOpts)
		}
		SetReaderIteratorPool(iteratorPool).