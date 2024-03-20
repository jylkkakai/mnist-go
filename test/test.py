import numpy as np
import tensorflow as tf

h = 20
w = 30
c = 32

# in_data = np.random.randint(-128, 127, size=(h, w, c))
in_data = np.random.randint(-128, 127, size=(1, 16))
in_data = np.array((0.05, 0.1))

model = tf.keras.Sequential(
    [tf.keras.layers.Dense(2, input_shape=(16,)), activation="sigmoid"]
    [tf.keras.layers.Dense(2, input_shape=(16,)), activation="sigmoid"]
)
model.compile(
    optimizer="adam",
    loss=tf.keras.losses.SparseCategoricalCrossentropy(from_logits=True),
    metrics=["accuracy"],
)
wgts = model.get_weights()
# print(type(wgts[0]))
print(wgts[0].shape)
# wgt = wgts[0].transpose((1, 0))
print(wgt.shape)
# print(wgt.shape)
out = model.predict(in_data)
# print(type(out))
# print(out.shape)

print(in_data.shape)
print(wgt)
print(out.shape)

#
# def write_flat(filename, arr):
#     with open(filename, "w") as outfile:
#         for i in arr:
#             for data in i:
#                 outfile.write(str(data) + " ")
#             # outfile.write("\n")


# write_flat("inpy.txt", in_data)
# write_flat("wgtpy.txt", wgt)
# write_flat("outpy.txt", out)
