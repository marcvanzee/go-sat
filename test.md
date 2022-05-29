# JAX Converters Evaluation Results

*Last generated on: 2022-05-29* (YYYY-MM-DD)

This file contains the evaluation results for all converters in table format.
Please see [README.md](README.md) for more details.

## Summary Table

| Example | jax2tf_to_tfjs |
| --- | --- |
| flax/transformer_lm1b | [NO](#error-trace-modelflaxtransformer_lm1b-converterjax2tf_to_tfjs) |

## Errors

## Error trace: model=flax/transformer_lm1b, converter=jax2tf_to_tfjs
```
ValueError("in user code:


    ValueError: Got a non-Tensor value FrozenDict({
        cache: {
            decoder: {
                encoderdecoderblock_0: {
                    SelfAttention_0: {
                        cache_index: <tf.Tensor 'StatefulPartitionedCall:1' shape=() dtype=int32>,
                        cached_key: <tf.Tensor 'StatefulPartitionedCall:2' shape=(2, 1, 1, 2) dtype=float32>,
                        cached_value: <tf.Tensor 'StatefulPartitionedCall:3' shape=(2, 1, 1, 2) dtype=float32>,
                    },
                },
                posembed_output: {
                    cache_index: <tf.Tensor 'StatefulPartitionedCall:4' shape=() dtype=uint32>,
                },
            },
        },
    }) for key 'output_1' in the output of the function __inference_tf_graph_1102 used to generate the SavedModel signature 'serving_default'. Outputs for functions used as signatures must be a single Tensor, a sequence of Tensors, or a dictionary from string to Tensor.
")
```
[Back to top](#summary-table)

See `converters_eval.py` for instructions on how to regenerate this table.
