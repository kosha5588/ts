"use strict";

// Get the new year
const getNewYear = () => {
  const currentYear = new Date().getFullYear();
  return new Date(`December 25 ${currentYear} 00:00:00`);
};

// update the year element
const year = document.querySelector(".year");
year.innerHTML = getNewYear().getFullYear();

// select elements
const app = document.querySelector(".countdown-timer");
const message = document.querySelector(".message");
const heading = document.querySelector("h1");

const format = (t) => {
  return t < 10 ? "0" + t : t;
};

const render = (time) => {
  app.innerHTML = `
        <div class="count-down">
            <div class="timer">
                <h3 class="days">${format(time.days)}</h3>
                <small>Days</small>
            </div>
            <div class="timer">
                <h3 class="hours">${format(time.hours)}</h3>
                <small>Hours</small>
            </div>
            <div class="timer">
                <h3 class="minutes">${format(time.minutes)}</h3>
                <small>Minutes</small>
            </div>
            <div class="timer">
                <h3 class="seconds">${format(time.seconds)}</h3>
                <small>Seconds</small>
            </div>
        </div>
        `;
};

const showMessage = () => {
  message.innerHTML = `Cristmas ${newYear}!`;
  app.innerHTML = "";
  heading.style.display = "none";
};

const hideMessage = () => {
  message.innerHTML = "";
  heading.style.display = "block";
};

const complete = () => {
  showMessage();

  // restart the countdown after showing the
  // greeting message for a day ()
  setTimeout(() => {
    hideMessage();
    countdownTimer.setExpiredDate(getNewYear());
  }, 1000 * 60 * 60 * 24);
};

const countdownTimer = new CountDown(getNewYear(), render, complete);
