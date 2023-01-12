use notify::{Config, RecommendedWatcher, RecursiveMode, Watcher};
use std::path::Path;
use std::time::Duration;
use std::thread;

fn main() {
    let path = std::env::args()
        .nth(1)
        .expect("Argument 1 needs to be a path");
    println!("watching {}", path);
    if let Err(e) = watch(path) {
        println!("error: {:?}", e)
    }
}

fn watch<P: AsRef<Path>>(path: P) -> notify::Result<()> {
    let (watcher_tx, watcher_rx) = crossbeam_channel::unbounded();
    let (sig_tx, sig_rx) = crossbeam_channel::unbounded();

    let mut watcher = RecommendedWatcher::new(watcher_tx, Config::default())?;
    watcher.watch(path.as_ref(), RecursiveMode::Recursive)?;

    ctrlc::set_handler(move || sig_tx.send(()).expect("could not send signal on channel."))
        .expect("error setting Ctrl-C handler");

    let th = thread::spawn(move || loop {
        crossbeam_channel::select! {
            recv(watcher_rx) -> res => {
              match res {
                  Ok(event) => println!("changed: {:?}", event),
                  Err(e) => println!("watch error: {:?}", e),
              }
            }
            recv(sig_rx) -> _res => {
                println!("got signal");
                break
            }
            default(Duration::from_secs(10)) => {
                println!("timeout");
                break
            }
        }
    });

    th.join().unwrap();
    println!("watch end");
    Ok(())
}
