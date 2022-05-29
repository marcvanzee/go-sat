# JAX Converters Evaluation Results

*Last generated on: 2022-05-29* (YYYY-MM-DD)

This file contains the evaluation results for all converters in table format.
Please see [README.md](README.md) for more details.

## Summary Table

| Example | jax2tf_xla | jax2tf_to_tfjs | jax2tf_to_tflite |
| --- | --- | --- | --- |
| flax/actor_critic | YES | YES | YES |
| flax/bilstm | YES | YES | [NO](#Example:-flax/bilstm,-Converter:-jax2tf_to_tflite) | 
| flax/resnet50 | YES | YES | YES |
| flax/transformer_lm1b | [NO](#Example:-flax/transformer_lm1b,-Converter:-jax2tf_xla) |  [NO](#Example:-flax/transformer_lm1b,-Converter:-jax2tf_to_tfjs) |  [NO](#Example:-flax/transformer_lm1b,-Converter:-jax2tf_to_tflite) | 
| flax/cnn | YES | YES | YES |

## Errors

## Example: flax/bilstm, Converter: jax2tf_to_tflite
```
RuntimeError('third_party/tensorflow/lite/kernels/concatenation.cc:158 t->dims->data[d] != t0->dims->data[d] (3 != 1)Node number 11 (CONCATENATION) failed to prepare.Node number 29 (WHILE) failed to invoke.')
```
## Error trace: model=flax/bilstm, converter=jax2tf_to_tflite)
```
InvalidArgumentError()
```
## Example: flax/transformer_lm1b, Converter: jax2tf_to_tfjs
```
ValueError("in user code:\n\n\n    ValueError: Got a non-Tensor value FrozenDict({\n        cache: {\n            decoder: {\n                encoderdecoderblock_0: {\n                    SelfAttention_0: {\n                        cache_index: <tf.Tensor 'StatefulPartitionedCall:1' shape=() dtype=int32>,\n                        cached_key: <tf.Tensor 'StatefulPartitionedCall:2' shape=(2, 1, 1, 2) dtype=float32>,\n                        cached_value: <tf.Tensor 'StatefulPartitionedCall:3' shape=(2, 1, 1, 2) dtype=float32>,\n                    },\n                },\n                posembed_output: {\n                    cache_index: <tf.Tensor 'StatefulPartitionedCall:4' shape=() dtype=uint32>,\n                },\n            },\n        },\n    }) for key 'output_1' in the output of the function __inference_tf_graph_183860 used to generate the SavedModel signature 'serving_default'. Outputs for functions used as signatures must be a single Tensor, a sequence of Tensors, or a dictionary from string to Tensor.\n")
```
## Example: flax/transformer_lm1b, Converter: jax2tf_to_tflite
```
TypeError("The DType <class 'numpy._FloatAbstractDType'> could not be promoted by <class 'numpy.dtype[str_]'>. This means that no common DType exists for the given inputs. For example they cannot be stored in a single array unless the dtype is `object`. The full list of DTypes is: (<class 'numpy.dtype[str_]'>, <class 'numpy._FloatAbstractDType'>)")
```

See `converters_eval.py` for instructions on how to regenerate this table.
