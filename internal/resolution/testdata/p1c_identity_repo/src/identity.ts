interface Shared {}
const Shared = () => 1;

function firstClock() {
  const time = Date.now();
  return time;
}

function secondClock() {
  const time = Date.now();
  return time;
}

function firstReport() {
  const now = Date.now();
  return now;
}

function secondReport() {
  const now = Date.now();
  return now;
}
