import numpy as np
import tensorflow as tf

h = 20
w = 30
c = 32


in_data = np.array([[0.05, 0.1]])
print(in_data.shape)
target = np.array([(0.01, 0.99)]).reshape(1, 2)


input = tf.keras.layers.Input(shape=(2,))
layer1 = tf.keras.layers.Dense(2, activation="sigmoid")(input)
output = tf.keras.layers.Dense(2, activation="sigmoid")(layer1)
model = tf.keras.Model(input, output)
model.compile(
    optimizer=tf.keras.optimizers.SGD(learning_rate=0.5),
    loss="mse",
    metrics=["accuracy"],
)
wgts = model.get_weights()
print(len(wgts))
print(wgts[0].shape)
wgts[0] = np.array([[0.15, 0.25], [0.2, 0.3]])
wgts[1] = np.array([0.35, 0.35])
wgts[2] = np.array([[0.4, 0.5], [0.45, 0.55]])
wgts[3] = np.array([0.6, 0.6])
print(len(wgts))
print(wgts[0].shape)
print(wgts)
model.set_weights(wgts)
print(wgts[0])
# np.save("../test_data/w1.npy", wgts[0].astype(np.float32))
# np.save("../test_data/b1.npy", wgts[1].astype(np.float32))
# np.save("../test_data/w2.npy", wgts[2].astype(np.float32))
# np.save("../test_data/b2.npy", wgts[3].astype(np.float32))
# np.save("../test_data/in.npy", in_data[0].astype(np.float32))
out = model.predict(in_data)
# np.save("../test_data/out.npy", out[0].astype(np.float32))
print(in_data.shape)
print(out.shape)
print(out)
loss = model.evaluate(in_data, target)
print(loss)

model.fit(in_data, target)
wgts = model.get_weights()
print(len(wgts))
print(wgts[0].shape)
igts = model.get_weights()
print(len(wgts))
print(wgts[0].shape)
print(wgts)
print(wgts[0].transpose(1, 0))
print(wgts[2].transpose(1, 0))
